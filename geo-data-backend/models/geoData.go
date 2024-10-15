package models

import (
    "gorm.io/gorm"
)

type GeoData struct {
    gorm.Model
    Name      string `json:"name"`
    GeoJSON   string `json:"geojson"`
    KML       string `json:"kml"`
    UserID    uint   `json:"user_id"`
}
