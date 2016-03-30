package goslackbot

import "time"

// SlackUser defines the go struct equivalent of the Slack RTM user:
// https://api.slack.com/types/user
type SlackUser struct {
	ID             string           `json:"id"`
	Name           string           `json:"name"`
	Deleted        bool             `json:"deleted"`
	Profile        SlackUserProfile `json:"profile"`
	IsAdmin        bool             `json:"is_admin"`
	IsOwner        bool             `json:"is_owner"`
	IsPrimaryOwner bool             `json:"is_primary_owner"`
	IsRestricted   bool             `json:"is_restricted"`
	Has2FA         bool             `json:"has_2fa"`
}

type SlackUserProfile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	RealName  string `json:"real_name"`
	Email     string `json:"email"`
	Skype     string `json:"skype"`
	Phone     string `json:"phone"`
	Image24   string `json:"image_24"`
}

type SlackChannel struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	IsArchived bool         `json:"is_archived"`
	IsMember   bool         `json:"is_member"`  // am i a member of this channel?
	IsGeneral  bool         `json:"is_general"` // The #general channel?
	Members    []string     `json:"members"`
	Latest     SlackMessage `json:"latest"`
}

type SlackPrivateChannel struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	IsArchived bool         `json:"is_archived"`
	IsMember   bool         `json:"is_member"`  // am i a member of this channel?
	IsGeneral  bool         `json:"is_general"` // The #general channel?
	Members    []string     `json:"members"`
	IsPrivate  bool         `json:"is_group"` // private channel?
	IsMPIM     bool         `json:"is_mpim"`
	Latest     SlackMessage `json:"latest"`
}

type SlackMPIM struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Members []string `json:"members"`
	Creator string   `json:"creator"`
}

type SlackIM struct {
	ID   string `json:"id"`
	User string `json:"user"`
}

// SlackRTMResponse the response we receive back from slack rtm.start
// as defined here: https://api.slack.com/methods/rtm.start
type SlackRTMResponse struct {
	Ok       bool                  `json:"ok"`
	Error    string                `json:"error"`
	Url      string                `json:"url"`
	Self     SlackRTMResponseSelf  `json:"self"`
	Users    []SlackUser           `json:"users"`
	Channels []SlackChannel        `json:"channels"`
	IMs      []SlackIM             `json:"ims"`
	MPIMs    []SlackMPIM           `jsonL:"mpims"`
	Groups   []SlackPrivateChannel `json:"groups"`
	Team     SlackTeam             `json:"team"`
}

type SlackRTMResponseSelf struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SlackMessage struct {
	Id        uint64           `json:"id"`
	Type      string           `json:"type"`
	SubType   string           `json:"sub_type"`
	Channel   string           `json:"channel"`
	Text      string           `json:"text"`
	User      string           `json:"user"`
	ReplyTo   uint64           `json:"reply_to, omitempty"`
	TimeStamp string           `json:"ts, omitempty"`
	Item      SlackMessageItem `json:"item, omitempty"`
	Name      string           `json:"name, omitempty"`
	Reaction  string           `json:"reaction, omitempty"`
}

type SlackMessageItem struct {
	Type      string `json:"type"`
	Channel   string `json:"channel"`
	TimeStamp string `json:"ts"`
}

type SlackTeam struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	EmailDomain string `json:"email_domain"`
	Domain      string `json:"domain"`
}

type SlackAPIReactionAdd struct {
	Token     string `json:"token"`
	Name      string `json:"name"`
	Channel   string `json:"channel"`
	TimeStamp string `json:"timestamp"`
}

type SlackPostMessageResponse struct {
	Ok        bool   `json:"ok"`
	Channel   string `json:"channel"`
	TimeStamp string `json:"ts"`
}

type SlackConversation struct {
	Messages []SlackMessage
	Ongoing  bool
	State    string
	Started  time.Time
}

type ConversationMap map[string]SlackConversation
