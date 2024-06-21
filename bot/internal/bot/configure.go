package bot

import "time"

func (b *Bot) SetName(name string) *Bot {
	b.name = name
	return b
}

func (b *Bot) Name() string {
	return b.name
}

func (b *Bot) SetPrimaryDomain(domain string) *Bot {
	b.domain = domain
	return b
}

func (b *Bot) PrimaryDomain() string {
	return b.domain
}

func (b *Bot) SetAllowedDomains(domains []string) *Bot {
	b.allowedDomains = domains
	return b
}

func (b *Bot) AddAllowedDomains(domain []string) *Bot {
	b.allowedDomains = append(b.allowedDomains, domain...)
	return b
}

func (b *Bot) AllowedDomains() []string {
	return b.allowedDomains
}

func (b *Bot) SetInterval(interval time.Duration) *Bot {
	b.interval = interval
	return b
}

func (b *Bot) Interval() time.Duration {
	return b.interval
}
