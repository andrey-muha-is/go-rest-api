package services

import (
	m "../models"
)

// GetChannels - return list of all channels
func GetChannels() []m.Channel {
	channels := []m.Channel{}

	// TODO: Add logic of retrieving channels list

	return channels
}

// GetChannelPrograms - return list of programs that belong to channel
func GetChannelPrograms(channelID  string) []m.Program {
	programs := []m.Program{}

	// TODO: Add logic of retrieving channel programs

	return programs
}
