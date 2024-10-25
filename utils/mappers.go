package utils

import (
	. "github.com/ivegotanidea/golang-gorm-postgres/models"
)

func MapBodyArts(bodyArts []ProfileBodyArt) []ProfileBodyArtResponse {
	bodyArtResponses := make([]ProfileBodyArtResponse, len(bodyArts))
	for i, bodyArt := range bodyArts {
		bodyArtResponses[i] = ProfileBodyArtResponse{
			ProfileID: bodyArt.ProfileID.String(),
			BodyArtID: bodyArt.BodyArtID,
		}
	}
	return bodyArtResponses
}

func MapPhotos(photos []Photo) []PhotoResponse {
	photoResponses := make([]PhotoResponse, len(photos))
	for i, photo := range photos {
		photoResponses[i] = PhotoResponse{
			URL:      photo.URL,
			Disabled: photo.Disabled,
			Approved: photo.Approved,
			Deleted:  photo.Deleted,
		}
	}
	return photoResponses
}

func MapProfileOptions(options []ProfileOption) []ProfileOptionResponse {
	optionResponses := make([]ProfileOptionResponse, len(options))
	for i, option := range options {
		optionResponses[i] = ProfileOptionResponse{
			Price:   option.Price,
			Comment: option.Comment,
			ProfileTag: ProfileTagResponse{
				Name: option.ProfileTag.Name,
			},
		}
	}
	return optionResponses
}

func MapService(service Service) *ServiceResponse {
	serviceResponse := ServiceResponse{
		ID:                   service.ID,
		ClientUserID:         service.ClientUserID,
		ClientUserRatingID:   service.ClientUserRatingID,
		ClientUserRating:     MapUserRating(service.ClientUserRating),
		ProfileID:            service.ProfileID,
		ProfileOwnerID:       service.ProfileOwnerID,
		ProfileRatingID:      service.ProfileRatingID,
		ProfileRating:        MapProfileRating(service.ProfileRating),
		DistanceBetweenUsers: service.DistanceBetweenUsers,
		TrustedDistance:      service.TrustedDistance,
		CreatedAt:            service.CreatedAt,
		UpdatedAt:            service.UpdatedAt,
		UpdatedBy:            service.UpdatedBy,
	}

	return &serviceResponse
}

func MapServices(services []Service) []ServiceResponse {
	serviceResponses := make([]ServiceResponse, len(services))
	for i, service := range services {
		serviceResponses[i] = *MapService(service)
	}
	return serviceResponses
}

func MapProfile(newProfile *Profile) *ProfileResponse {
	profileResponse := &ProfileResponse{
		ID:                     newProfile.ID.String(),
		UpdatedBy:              &newProfile.UpdatedBy,
		UserID:                 newProfile.UserID.String(),
		CityID:                 &newProfile.CityID,
		Active:                 newProfile.Active,
		Phone:                  newProfile.Phone,
		Name:                   newProfile.Name,
		Bio:                    newProfile.Bio,
		Age:                    newProfile.Age,
		Height:                 newProfile.Height,
		Weight:                 newProfile.Weight,
		Bust:                   newProfile.Bust,
		AddressLatitude:        newProfile.AddressLatitude,
		AddressLongitude:       newProfile.AddressLongitude,
		BodyTypeID:             newProfile.BodyTypeID,
		EthnosID:               newProfile.EthnosID,
		HairColorID:            newProfile.HairColorID,
		IntimateHairCutID:      newProfile.IntimateHairCutID,
		PrinceSaunaNightRatio:  newProfile.PrinceSaunaNightRatio,
		PriceCarNightRatio:     newProfile.PriceCarNightRatio,
		PriceVisitNightRatio:   newProfile.PriceVisitNightRatio,
		PriceInHouseNightRatio: newProfile.PriceInHouseNightRatio,
		PriceSaunaHour:         newProfile.PriceSaunaHour,
		PriceVisitHour:         newProfile.PriceVisitHour,
		PriceCarContact:        newProfile.PriceCarContact,
		PriceCarHour:           newProfile.PriceCarHour,
		PriceSaunaContact:      newProfile.PriceSaunaContact,
		PriceVisitContact:      newProfile.PriceVisitContact,
		PriceInHouseHour:       newProfile.PriceInHouseHour,
		PriceInHouseContact:    newProfile.PriceInHouseContact,
		ContactPhone:           newProfile.ContactPhone,
		ContactWA:              newProfile.ContactWA,
		ContactTG:              newProfile.ContactTG,
		Moderated:              newProfile.Moderated,
		ModeratedAt:            &newProfile.ModeratedAt,
		Verified:               newProfile.Verified,
		VerifiedAt:             &newProfile.VerifiedAt,
		VerifiedBy:             &newProfile.VerifiedBy,
		CreatedAt:              newProfile.CreatedAt,
	}

	profileResponse.BodyArts = MapBodyArts(newProfile.BodyArts)
	profileResponse.Photos = MapPhotos(newProfile.Photos)
	profileResponse.ProfileOptions = MapProfileOptions(newProfile.ProfileOptions)
	profileResponse.Services = MapServices(newProfile.Services)

	return profileResponse
}

func MapUserRating(userRating *UserRating) *UserRatingResponse {
	if userRating == nil {
		return nil
	}

	ratedUserTags := make([]RatedUserTagResponse, len(userRating.RatedUserTags))
	for i, tag := range userRating.RatedUserTags {
		ratedUserTags[i] = RatedUserTagResponse{
			Type: tag.Type,
			UserTag: UserTagResponse{
				Name: tag.UserTag.Name,
			},
		}
	}

	return &UserRatingResponse{
		ID:                userRating.ID,
		ServiceID:         userRating.ServiceID,
		UserID:            userRating.UserID,
		ReviewTextVisible: userRating.ReviewTextVisible,
		Review:            userRating.Review,
		Score:             userRating.Score,
		CreatedAt:         userRating.CreatedAt,
		UpdatedAt:         userRating.UpdatedAt,
		RatedUserTags:     ratedUserTags,
		UpdatedBy:         userRating.UpdatedBy,
	}
}

func MapProfileRating(profileRating *ProfileRating) *ProfileRatingResponse {
	if profileRating == nil {
		return nil
	}

	ratedProfileTags := make([]RatedProfileTagResponse, len(profileRating.RatedProfileTags))
	for i, tag := range profileRating.RatedProfileTags {
		ratedProfileTags[i] = RatedProfileTagResponse{
			Type: tag.Type,
			ProfileTag: ProfileTagResponse{
				Name: tag.ProfileTag.Name,
			},
		}
	}

	return &ProfileRatingResponse{
		ID:                profileRating.ID,
		ServiceID:         profileRating.ServiceID,
		ProfileID:         profileRating.ProfileID,
		ReviewTextVisible: profileRating.ReviewTextVisible,
		Review:            profileRating.Review,
		Score:             profileRating.Score,
		CreatedAt:         profileRating.CreatedAt,
		UpdatedAt:         profileRating.UpdatedAt,
		RatedProfileTags:  ratedProfileTags,
		UpdatedBy:         profileRating.UpdatedBy,
	}
}
