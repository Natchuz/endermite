package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type VersionManifest struct {
	Latest struct {
		Release  string `json:"release"`
		Snapshot string `json:"snapshot"`
	} `json:"latest"`
	Versions []VersionEntry `json:"versions"`
}

type VersionEntry struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	Time        string `json:"time"`
	ReleaseTime string `json:"releaseTime"`
}

func requestVersionManifest() (VersionManifest, error) {
	resp, err := http.Get("https://launchermeta.mojang.com/mc/game/version_manifest.json")
	if err != nil {
		return VersionManifest{}, err
	}
	defer resp.Body.Close()

	jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return VersionManifest{}, err
	}

	var manifest = VersionManifest{}
	err = json.Unmarshal(jsonDataFromHttp, &manifest)
	if err != nil {
		return VersionManifest{}, err
	}

	return manifest, nil
}
