package thirdclient

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"googlemaps.github.io/maps"
)

var (
	PlacesDetailFields = symmetricDifference(PlacesDetailFieldsBasicList, PlacesDetailFieldsContactList, PlacesDetailFieldsAtmosphereList)
	GoogleMapsKey      = "YOUR_GOOGLE_MAPS_API_KEY" // todo 请替换为你的 API Key
)

// 基本字段列表
var PlacesDetailFieldsBasicList = []maps.PlaceDetailsFieldMask{
	"address_component",
	"adr_address",
	"business_status",
	"formatted_address",
	"geometry",
	"geometry/location",
	"geometry/location/lat",
	"geometry/location/lng",
	"geometry/viewport",
	"geometry/viewport/northeast",
	"geometry/viewport/northeast/lat",
	"geometry/viewport/northeast/lng",
	"geometry/viewport/southwest",
	"geometry/viewport/southwest/lat",
	"geometry/viewport/southwest/lng",
	"icon",
	"name",
	"permanently_closed",
	"photo",
	"place_id",
	"plus_code",
	"type",
	"url",
	"utc_offset",
	"vicinity",
}

// 联系字段列表
var PlacesDetailFieldsContactList = []maps.PlaceDetailsFieldMask{
	"formatted_phone_number",
	"international_phone_number",
	"opening_hours",
	"website",
}

// 氛围字段列表
var PlacesDetailFieldsAtmosphereList = []maps.PlaceDetailsFieldMask{
	"price_level",
	"rating",
	"review",
	"user_ratings_total",
}

// symmetricDifference 返回多个字符串切片的对称差集（即出现奇数次的元素）
func symmetricDifference(lists ...[]maps.PlaceDetailsFieldMask) []maps.PlaceDetailsFieldMask {
	counts := make(map[maps.PlaceDetailsFieldMask]int)
	// 统计每个元素出现的次数
	for _, list := range lists {
		for _, item := range list {
			counts[item]++
		}
	}
	// 只保留出现奇数次的元素
	var result []maps.PlaceDetailsFieldMask
	for item, count := range counts {
		if count%2 == 1 {
			result = append(result, item)
		}
	}
	return result
}

func GetPlaceDetail(placeID string) (map[string]interface{}, error) {
	// 初始化 Google Maps 客户端
	// todo:是否每次都需要创建对象
	client, err := maps.NewClient(maps.WithAPIKey(GoogleMapsKey))
	if err != nil {
		log.Printf("Error creating Google Maps client: %v", err)
		return nil, fmt.Errorf("failed to create maps client")
	}

	req := &maps.PlaceDetailsRequest{
		PlaceID: placeID,
		Fields:  PlacesDetailFields,
	}

	resp, err := client.PlaceDetails(context.Background(), req)
	if err != nil {
		log.Printf("Error calling PlaceDetails: %v", err)
		return nil, fmt.Errorf("failed to get place details")
	}

	data, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
