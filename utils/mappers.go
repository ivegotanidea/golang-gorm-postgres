package utils

import (
	. "github.com/ivegotanidea/golang-gorm-postgres/models"
)

func mapBodyArts(bodyArts []ProfileBodyArt) []ProfileBodyArtResponse {
	bodyArtResponses := make([]ProfileBodyArtResponse, len(bodyArts))
	for i, bodyArt := range bodyArts {
		bodyArtResponses[i] = ProfileBodyArtResponse{
			ProfileID: bodyArt.ProfileID.String(),
			BodyArtID: bodyArt.BodyArtID,
		}
	}
	return bodyArtResponses
}

func mapPhotos(photos []Photo) []PhotoResponse {
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

func mapProfileOptions(options []ProfileOption) []ProfileOptionResponse {
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

func mapServices(services []Service) []ServiceResponse {
	serviceResponses := make([]ServiceResponse, len(services))
	for i, service := range services {
		serviceResponses[i] = ServiceResponse{
			ID:                   service.ID,
			ClientUserID:         service.ClientUserID,
			ClientUserRatingID:   service.ClientUserRatingID,
			ClientUserRating:     mapUserRating(service.ClientUserRating),
			ProfileID:            service.ProfileID,
			ProfileOwnerID:       service.ProfileOwnerID,
			ProfileRatingID:      service.ProfileRatingID,
			ProfileRating:        mapProfileRating(service.ProfileRating),
			DistanceBetweenUsers: service.DistanceBetweenUsers,
			TrustedDistance:      service.TrustedDistance,
			CreatedAt:            service.CreatedAt,
			UpdatedAt:            service.UpdatedAt,
			UpdatedBy:            service.UpdatedBy,
		}
	}
	return serviceResponses
}

func MapProfile(newProfile *Profile) *ProfileResponse {
	profileResponse := &ProfileResponse{
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

	profileResponse.BodyArts = mapBodyArts(newProfile.BodyArts)
	profileResponse.Photos = mapPhotos(newProfile.Photos)
	profileResponse.ProfileOptions = mapProfileOptions(newProfile.ProfileOptions)
	profileResponse.Services = mapServices(newProfile.Services)

	return profileResponse
}

func mapUserRating(userRating *UserRating) *UserRatingResponse {
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

func mapProfileRating(profileRating *ProfileRating) *ProfileRatingResponse {
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
