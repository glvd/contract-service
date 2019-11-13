package main

import (
	"github.com/glvd/conversion"
	"github.com/glvd/seed/model"
)

// SeedVideoToConversionVideo ...
func SeedVideoToConversionVideo(v model.Video) conversion.Video {
	vd := conversion.Video{
		No:           v.Bangumi,
		Intro:        v.Intro,
		Alias:        v.Alias,
		ThumbHash:    v.ThumbHash,
		PosterHash:   v.PosterHash,
		SourceHash:   v.SourceHash,
		M3U8Hash:     v.M3U8Hash,
		Key:          v.Key,
		M3U8:         v.M3U8,
		Role:         v.Role,
		Director:     v.Director,
		Systematics:  v.Systematics,
		Season:       v.Season,
		TotalEpisode: v.TotalEpisode,
		Episode:      v.Episode,
		Producer:     v.Producer,
		Publisher:    v.Publisher,
		Type:         v.Type,
		Format:       v.Format,
		Language:     v.Language,
		Caption:      v.Caption,
		Group:        v.Group,
		Index:        v.Index,
		Date:         v.Date,
		Sharpness:    v.Sharpness,
		Series:       v.Series,
		Tags:         v.Tags,
		Length:       v.Length,
		Sample:       []string{},
		Uncensored:   v.Uncensored,
	}
	vd.SetVersion(v.Version)
	vd.SetID(v.ID)
	return vd
}
