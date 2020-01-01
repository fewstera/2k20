package nyecanvas

import (
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
)

type Screen int

const (
	CurrentTime Screen = iota
	SecRemaining
	MinSecRemaining
	DisplayImage
	ScrollingText
)

const timestampFile string = "./last_message_ts"

type MessagesResponse struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	Message   string `json:"message"`
	From      string `json:"from"`
	Timestamp uint64 `json:"timestamp"`
}

type Canvas struct {
	width              int
	height             int
	smallFont          font.Face
	bigFont            font.Face
	massiveFont        font.Face
	displayImage       image.Image
	screenStartTime    time.Time
	midnight           time.Time
	textDisplayEndTime time.Time
	messages           []Message
	textToDisplay      string
}

func New(width, height int) (*Canvas, error) {
	smallFont, err := loadFont("6x10")
	if err != nil {
		return nil, fmt.Errorf("loading small font: %s", err)
	}

	bigFont, err := loadFont("6x12")
	if err != nil {
		return nil, fmt.Errorf("loading big font: %s", err)
	}

	massiveFont, err := loadFont("10x20")
	if err != nil {
		return nil, fmt.Errorf("loading massive font: %s", err)
	}

	displayImage, err := gg.LoadPNG("./images/tree.png")
	if err != nil {
		return nil, fmt.Errorf("loading displayImage: %s", err)
	}

	c := &Canvas{
		width:           width,
		height:          height,
		smallFont:       smallFont,
		bigFont:         bigFont,
		massiveFont:     massiveFont,
		displayImage:    displayImage,
		screenStartTime: time.Now(),
		midnight:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	go c.fetchMessagesLoop()

	return c, nil
}

func (c *Canvas) Tick() image.Image {
	dc := gg.NewContext(c.width, c.height)
	dc.SetRGB(0, 0, 0)
	dc.DrawRectangle(0, 0, float64(c.width), float64(c.height))
	dc.Fill()

	switch c.CurrentScreen() {
	case CurrentTime:
		c.DrawCurrentTimeScreen(dc)
	case MinSecRemaining:
		c.DrawMinSecRemaining(dc)
	case DisplayImage:
		c.DrawDisplayImage(dc)
	case SecRemaining:
		c.DrawSecRemaining(dc)
	case ScrollingText:
		c.DrawScrollingTextScreen(dc)
	}

	return dc.Image()
}

func (c *Canvas) CurrentScreen() Screen {
	now := time.Now()
	ttm := int(c.midnight.Sub(now).Seconds())
	advertText := "Submit your msgs at https://nye.fewstera.com"

	if ttm < 99 && ttm >= 0 {
		return SecRemaining
	}

	if ttm < 0 && ttm >= -120 {
		if c.textToDisplay != "Happy New Year!" {
			c.StartDisplayingText("Happy New Year!")
		}
		return ScrollingText
	}

	messageIsDisplaying := now.Sub(c.textDisplayEndTime).Seconds() <= 0

	if len(c.messages) > 0 {
		if !messageIsDisplaying || c.textToDisplay == advertText {
			msgs := c.messages
			newMessage, msgs := c.messages[0], c.messages[1:]
			c.messages = msgs
			c.StartDisplayingText(fmt.Sprintf("Msg from %s: %s", newMessage.From, newMessage.Message))
			return ScrollingText
		}
	}

	if messageIsDisplaying {
		return ScrollingText
	}

	screens := []Screen{DisplayImage, CurrentTime}
	if ttm > 0 {
		screens = append(screens, MinSecRemaining)
	}

	choice := int((time.Now().Unix() / 30)) % (len(screens) + 1)
	if choice == len(screens) {
		c.StartDisplayingText(advertText)
		return ScrollingText
	}
	return screens[choice]
}

func (c *Canvas) fetchMessagesLoop() {
	for {
		c.fetchMessages()
		time.Sleep(time.Duration(5) * time.Second)
	}
}

func (c *Canvas) fetchMessages() {
	lastReadTimestamp, err := ioutil.ReadFile(timestampFile)
	if err != nil {
		lastReadTimestamp = []byte("0")
	}

	res, err := http.Get(fmt.Sprintf("https://nye.fewstera.com/api/messages?since=%s", lastReadTimestamp))
	if err != nil {
		fmt.Printf("Error fetching messages: %s", err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading messages body: %s", err)
		return
	}

	messagesRes := MessagesResponse{}
	err = json.Unmarshal(body, &messagesRes)
	if err != nil {
		fmt.Printf("Error unmarshalling messages body: %s", err)
		return
	}

	// Continute if we have no new messages:
	if len(messagesRes.Messages) == 0 {
		return
	}

	lastMessageTimestamp := []byte(fmt.Sprintf("%d", messagesRes.Messages[len(messagesRes.Messages)-1].Timestamp))
	err = ioutil.WriteFile(timestampFile, lastMessageTimestamp, 0644)
	if err != nil {
		fmt.Printf("Error saving timestamp file: %s", err)
		return
	}

	c.messages = append(c.messages, messagesRes.Messages...)
}
