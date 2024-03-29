package discord

import (
	"fmt"
	"os"
	"time"

	embed "github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
	"github.com/gookit/slog"
	"github.com/keshon/discord-bot-template/internal/config"
	"github.com/keshon/discord-bot-template/internal/version"
	"github.com/keshon/discord-bot-template/mod-helloworld/utils"
)

// handleHelpCommand handles the help command for the Discord bot.
//
// Takes in a session and a message create, and does not return any value.
func (d *Discord) handleHelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	d.changeAvatar(s)

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Fatal("Error loading config:", err)
		return
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = cfg.RestHostname
	}

	avatarURL := utils.InferProtocolByPort(host, 443) + host + "/avatar/random?" + fmt.Sprint(time.Now().UnixNano())
	prefix := d.CommandPrefix

	hello := fmt.Sprintf("`%vhello` — print «Hello World»\n", prefix)
	hi := fmt.Sprintf("`%vhi` — print «Hello Galaxy»\n", prefix)
	help := fmt.Sprintf("`%vhelp`, `%vh` — show help\n", prefix, prefix)
	about := fmt.Sprintf("`%vabout`, `%va` — show version\n", prefix, prefix)
	register := fmt.Sprintf("`%vregister` — enable commands listening\n", prefix)
	unregister := fmt.Sprintf("`%vunregister` — disable commands listening", prefix)

	embedMsg := embed.NewEmbed().
		SetTitle(fmt.Sprintf("ℹ️ %v — Command Usage", version.AppName)).
		SetDescription("Some commands are aliased for shortness.\n").
		AddField("", "**Demo**\n"+hello+hi).
		AddField("", "").
		AddField("", "**Information**\n"+help+about).
		AddField("", "").
		AddField("", "**Managing Bot**\n"+register+unregister).
		SetThumbnail(avatarURL).
		SetColor(0x9f00d4).
		SetFooter(version.AppFullName).
		MessageEmbed

	_, err = s.ChannelMessageSendEmbed(m.Message.ChannelID, embedMsg)
	if err != nil {
		slog.Fatal("Error sending embed message", err)
	}
}
