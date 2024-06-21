package bot

import (
	"time"

	"github.com/gocolly/colly"
)

type Bot struct {
	collector      *colly.Collector
	target         string
	pattern        string
	domain         string
	allowedDomains []string
	interval       time.Duration
}

func New() *Bot {
	return &Bot{
		collector:      nil,
		target:         "",
		pattern:        "",
		domain:         "",
		allowedDomains: []string{},
		interval:       time.Minute * 10,
	}
}

func (b *Bot) Build() *Bot {
	b.collector = colly.NewCollector(
		colly.AllowedDomains(b.allowedDomains...),
	)

	return b
}
