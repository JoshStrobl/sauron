package main

import (
	"fmt"
	"github.com/JoshStrobl/trunk"
	"github.com/PuerkitoBio/goquery"
	"github.com/TryStreambits/sauron"
	"net/url"
	"strings"
)

func main() {
	image, imageLinkErr := sauron.GetLink("https://i3.ytimg.com/vi/OE-Y-PotqTQ/maxresdefault.jpg")

	if imageLinkErr == nil { // Got the Image
		if image.Extras["IsImageLink"] == "true" { // If we have an Extras metadata set
			trunk.LogSuccess("Got JPG")
			fmt.Printf("%v\n", image)
		} else {
			trunk.LogErr("Failed to determine that our image is precisely that.")
		}
	} else {
		trunk.LogErr(fmt.Sprintf("Failed to get Image: %v", imageLinkErr))
	}

	video, videoLinkErr := sauron.GetLink("http://mirrors.standaloneinstaller.com/video-sample/small.mp4")

	if videoLinkErr == nil { // Got the video
		if video.Extras["IsVideoLink"] == "true" { // If we have an Extras metadata set
			trunk.LogSuccess("Got MP4")
			fmt.Printf("%v\n", video)
		} else {
			trunk.LogErr("Failed to determine that our video is prcisely that.")
		}
	} else {
		trunk.LogErr(fmt.Sprintf("Failed to get Video: %v", videoLinkErr))
	}

	twitter, twitterLinkErr := sauron.GetLink("https://twitter.com/trystreambits/status/1246090584714027010")

	if twitterLinkErr == nil { // Got Twitter link data
		trunk.LogSuccess("Got @trystreambits Tweet")
		fmt.Printf("%v\n", twitter)
	} else {
		trunk.LogErr(fmt.Sprintf("Failed to get Tweet: %v", twitterLinkErr))
	}

	twitchStreamer, twitchStreamerLinkErr := sauron.GetLink("https://www.twitch.tv/towelliee")

	if twitchStreamerLinkErr == nil { // Successfully got the page
		if twitchStreamer.Extras["Streamer"] != "Towelliee" || // Streamer isn't Towelliee
			twitchStreamer.Extras["Game"] == "" || // Game is empty
			!strings.HasPrefix(twitchStreamer.Extras["GameLink"], "https://www.twitch.tv/directory/game/") || // Not expected beginning of URL for game directory listing
			!strings.HasPrefix(twitchStreamer.Extras["GameArtFull"], "https://static-cdn.jtvnw.net/ttv-boxart/") { // Not expected beginning of URL for box art
			trunk.LogErr(fmt.Sprintf("Fetched Streamer details but does not match expectation: %s", twitchStreamer))
		} else {
			trunk.LogSuccess(fmt.Sprintf("Got Twitch streamer details: %v", twitchStreamer))
		}
	} else {
		trunk.LogErr(fmt.Sprintf("Failed to get Twitch streamer: %v", twitchStreamerLinkErr))
	}

	twitchClip, twitchClipLinkErr := sauron.GetLink("https://www.twitch.tv/towelliee/clip/VastTentativeDinosaurGOWSkull")

	if twitchClipLinkErr == nil { // Got the clip
		if twitchClip.Title != "Towelliee - eclipse - Twitch" || // Title doesn't match expected
			twitchClip.Extras["ClipName"] != "eclipse" || // Not eclipse
			twitchClip.Extras["ClipSlug"] != "VastTentativeDinosaurGOWSkull" || // Slug doesn't match expected
			twitchClip.Extras["Game"] != "World of Warcraft" || // Game doesn't match expectation
			twitchClip.Extras["GameLink"] != "https://www.twitch.tv/directory/game/World of Warcraft" || // Not expected URL for game
			twitchClip.Extras["GameArtFull"] != "https://static-cdn.jtvnw.net/ttv-boxart/World%20of%20Warcraft.jpg" { // Full game art doesn't match
			trunk.LogErr(fmt.Sprintf("Fetched Clip details but does not match expectation: %v", twitchClip))
		} else {
			trunk.LogSuccess(fmt.Sprintf("Got Twitch clip details: %v", twitchClip))
		}
	} else { // Failed to get the clip
		trunk.LogErr(fmt.Sprintf("Failed to get the Twitch clip: %v", twitchClipLinkErr))
	}

	twitchSecondaryClip, twitchSecondaryClipLinkErr := sauron.GetLink("https://clips.twitch.tv/VastTentativeDinosaurGOWSkull")

	if twitchSecondaryClipLinkErr == nil { // Got the clip
		if twitchSecondaryClip.Title != twitchClip.Title { // If our Title doesn't match our proper full Twitch clip URL
			trunk.LogErr(fmt.Sprintf("Fetched Clip details but does not match expectation: %v", twitchSecondaryClip))
		} else {
			trunk.LogSuccess(fmt.Sprintf("Got Twitch clip details for clips.twitch.tv subdomain: %v", twitchSecondaryClip))
		}
	} else {
		trunk.LogErr(fmt.Sprintf("Failed to get the Twitch clip via clips.twitch.tv: %v", twitchSecondaryClipLinkErr))
	}

	bigBuckBunnyLink, linkErr := sauron.GetLink("https://www.youtube.com/watch?v=YE7VzlLtp-4")

	if linkErr == nil { // Successfully got link data
		if bigBuckBunnyLink.Title == "Big Buck Bunny" && bigBuckBunnyLink.Extras["IsVideo"] == "true" { // Successfully fetched
			trunk.LogSuccess(fmt.Sprintf("Fetched Big Buck Bunny. Has the following content: %s", bigBuckBunnyLink))
		} else { // Details do not match
			trunk.LogErr(fmt.Sprintf("Successfully fetched Big Buck Bunny but content does not match expectation: %s", bigBuckBunnyLink))
		}
	} else { // If we failed to fetch Big Buck Bunny
		trunk.LogErr(fmt.Sprintf("Failed to get Big Buck Bunny: %v", linkErr))
	}

	playlistTestLink, playlistTestLinkErr := sauron.GetLink("https://www.youtube.com/playlist?list=PLFF5D72E24079FB50")

	if playlistTestLinkErr == nil { // Successfully got playlist
		if playlistTestLink.Title == "Mat Kearney - Young Love" && // Name matches
			playlistTestLink.Extras["IsPlaylist"] == "true" && // Is a Playlist
			playlistTestLink.Image == "https://i.ytimg.com/vi/FANROVxej50/hqdefault.jpg" { // Playlist Image matches
			trunk.LogSuccess(fmt.Sprintf("Fetched Youtube Playlist. Has the following content: %s\n", playlistTestLink))
		} else {
			trunk.LogErr(fmt.Sprintf("Successfully fetched Youtube Playlist but content does not match expectation: %s\n", playlistTestLink))
		}
	} else {
		trunk.LogErr(fmt.Sprintf("Failed to get Youtube Playlist: %v", playlistTestLink))
	}

	redditPost, redditLinkErr := sauron.GetLink("https://www.reddit.com/r/SolusProject/comments/b2a8x0/solus_4_fortitude_released_solus/")

	if redditLinkErr == nil { // Successfully got reddit post
		if redditPost.Title == "Solus 4 Fortitude Released | Solus : SolusProject" && redditPost.Extras["Likes"] != "" { // Successfully got Reddit post
			trunk.LogSuccess(fmt.Sprintf("Fetched Reddit post. Has the following content: %s\n", redditPost))
		} else { // Failed to get reddit post, potentially likes
			trunk.LogErr(fmt.Sprintf("Successfully fetched Reddit post but content does not match expectations: %s\n", redditPost))
		}
	} else { // Failed to fetch Reddit post
		trunk.LogErr(fmt.Sprintf("Failed to get Reddit post: %v", redditLinkErr))
	}

	downvotedPost, redditDownvoteLinkErr := sauron.GetLink("https://old.reddit.com/r/linux/comments/ielvry/linux_used_to_be_to_bring_life_to_your_old/")

	if redditDownvoteLinkErr == nil { // Successfully got the downvoted reddit post
		trunk.LogSuccess(fmt.Sprintf("Fetched downvoted Reddit post. Has the following content: %s\n", downvotedPost))
	}
	sauron.Register("joshuastrobl.com", PersonalSiteHandler)

	personalSiteLink, personalLinkErr := sauron.GetLink("https://joshuastrobl.com")

	if personalLinkErr == nil { // Successfully got personal site
		if personalSiteLink.Title == "Home | Joshua Strobl" && strings.HasPrefix(personalSiteLink.Extras["Generator"], "Hugo") { // Successfully got Personal Site
			trunk.LogSuccess(fmt.Sprintf("Fetched Personal Site. Has the following content: %s\n", personalSiteLink))
		} else { // Failed to get personal site, potentially generator info
			trunk.LogErr(fmt.Sprintf("Successfully fetched Personal Site but content does not match expecations: %s\n", personalSiteLink))
		}
	} else { // Failed to get personal site
		trunk.LogErr(fmt.Sprintf("Failed to get Personal Site: %v", personalLinkErr))
	}

	gogLink, gogLinkErr := sauron.GetLink("https://www.gog.com/game/the_witcher")

	if gogLinkErr == nil { // Got GOG
		if strings.HasSuffix(gogLink.Title, "The Witcher: Enhanced Edition on GOG.com") { // If we successfully fetched the title when they reuse it weirdly
			trunk.LogSuccess(fmt.Sprintf("Fetched GOG site. Has the following content: %s\n", gogLink))
		} else { // Failed to get the correct title
			trunk.LogErr(fmt.Sprintf("Failed to fetch the GOG site which has weird title re-use: %s\n", gogLink))
		}
	} else {
		trunk.LogErr(fmt.Sprintf("Failed to get GOG: %v", personalLinkErr))
	}
}

// PersonalSiteHandler is a LinkParser for joshuastrobl.com
func PersonalSiteHandler(doc *goquery.Document, u *url.URL, fullPath string) (link *sauron.Link, parseErr error) {
	link, parseErr = sauron.Primitive(doc, u, fullPath) // Handle with Primitive first

	if parseErr != nil { // If we failed with primitive
		return
	}

	generatorElem := doc.Find(`meta[name="generator"]`) // Get the meta generator tag
	generator := generatorElem.AttrOr("content", "ERROR")
	link.Extras["Generator"] = generator

	return
}
