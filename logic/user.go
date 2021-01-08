package logic

import (
	"context"
	"fmt"

	"github.com/parijatpurohit/vaccinepass/server/transport"
	"github.com/parijatpurohit/vaccinepass/zz_generated/go/protogen"
)

func (l *logicImpl) GetUserSummary(req *transport.GetUserSummaryRequest) (*transport.GetUserSummaryResponse, error) {
	userDocs, err := l.client.UserDocs_FindByUserID(context.Background(), &protogen.UserDocs_FindByUserIDRequest{
		Query: &protogen.UserDocs_FindByUserID_Query{UserID: req.UserUUID},
	})
	if err != nil {
		return nil, err
	}

	var docIDs []string
	for _, doc := range userDocs.UserDocss {
		docIDs = append(docIDs, doc.UUID)
	}

	userVaccines, err := l.client.UserVaccines_FindByUserDocID(context.Background(), &protogen.UserVaccines_FindByUserDocIDRequest{
		Query: &protogen.UserVaccines_FindByUserDocID_Query{UserDocID: docIDs},
	})
	if err != nil {
		return nil, err
	}
	total := len(userVaccines.UserVacciness)

	var vaccineUUIDs []string
	for _, uv := range userVaccines.UserVacciness {
		vaccineUUIDs = append(vaccineUUIDs, uv.VaccineID)
	}

	vaccines, err := l.getVaccines(vaccineUUIDs)
	if err != nil {
		return nil, err
	}

	var latestVaccinations []*transport.VaccineDetails

	for index, userVaccine := range userVaccines.UserVacciness {
		latestVaccinations = append(latestVaccinations, &transport.VaccineDetails{
			Name:      vaccines[index].Name,
			VaccineID: userVaccine.VaccineID,
			UserID:    req.UserUUID,
			Authority: userVaccine.VaccineAuthorityID,
			Country:   userVaccine.VaccineAuthorityID,
		})
	}

	return &transport.GetUserSummaryResponse{
		Summary: &transport.UserSummary{TotalVaccinations: int64(total), LatestVaccinations: latestVaccinations},
	}, nil
}

func (l *logicImpl) GetUserVaccineDetails(req *transport.GetUserVaccineDetailsRequest) (*transport.GetUserVaccineDetailsResponse, error) {
	userDetails, err := l.client.UserDocs_FindByUserID(context.Background(), &protogen.UserDocs_FindByUserIDRequest{
		Query: &protogen.UserDocs_FindByUserID_Query{
			UserID: req.UserUUID,
		},
	})

	if err != nil {
		return nil, err
	}

	var userDocIDs []string
	userDocDetails := map[string]string{}
	for _, userDoc := range userDetails.UserDocss {
		userDocIDs = append(userDocIDs, userDoc.UUID)
		userDocDetails[userDoc.UUID] = userDoc.OfficialIDType
	}

	data, err := l.client.UserVaccines_FindByUserDocID(context.Background(), &protogen.UserVaccines_FindByUserDocIDRequest{
		Query: &protogen.UserVaccines_FindByUserDocID_Query{
			UserDocID: userDocIDs,
		},
	})
	if err != nil {
		return nil, err
	}

	var vaccineIDs []string
	for _, v := range data.UserVacciness {
		vaccineIDs = append(vaccineIDs, v.VaccineID)
	}

	vaccines, err := l.getVaccines(vaccineIDs)
	if err != nil {
		return nil, err
	}
	if len(vaccines) != len(vaccineIDs) {
		return nil, fmt.Errorf("invalid vaccine ids %v!= %v", len(vaccines), len(vaccineIDs))
	}

	var uv []*transport.VaccineDetails
	for index, row := range data.UserVacciness {
		uv = append(uv, &transport.VaccineDetails{
			Name:            vaccines[index].Name,
			VaccineID:       vaccines[index].UUID,
			UserID:          row.UserDocID,
			UserIDType:      userDocDetails[row.UserDocID],
			Authority:       row.VaccineAuthorityID,
			VaccinationDate: row.VaccinationDate,
			Certificate:     row.Certificate,
		})
	}
	return &transport.GetUserVaccineDetailsResponse{VaccineDetails: uv}, nil
}

func (l *logicImpl) getVaccines(vaccineIDs []string) ([]*protogen.Vaccines, error) {
	vaccines, err := l.client.Vaccines_FindByVaccineID(context.Background(), &protogen.Vaccines_FindByVaccineIDRequest{
		Query: &protogen.Vaccines_FindByVaccineID_Query{
			UUID: vaccineIDs,
		},
	})
	if err != nil {
		return nil, err
	}
	return vaccines.Vacciness, nil
}

func (l *logicImpl) GetCountryVaccines(req *transport.GetCountryVaccinesRequest) (*transport.GetCountryVaccinesResponse, error) {
	vaccines, err := l.client.CountryVaccines_FindByCountryID(context.Background(), &protogen.CountryVaccines_FindByCountryIDRequest{
		Query: &protogen.CountryVaccines_FindByCountryID_Query{
			CountryID: req.CountryID,
		},
	})

	if err != nil {
		return nil, err
	}

	var vaccineIDs []string
	for _, v := range vaccines.CountryVacciness {
		vaccineIDs = append(vaccineIDs, v.VaccineID)
	}

	vaccineData, err := l.getVaccines(vaccineIDs)

	if err != nil {
		return nil, err
	}

	var vd []*transport.RequiredVaccineDetails
	for _, row := range vaccineData {
		vd = append(vd, &transport.RequiredVaccineDetails{
			Name:        row.Name,
			Description: row.Name + row.UUID,
		})
	}

	res := &transport.GetCountryVaccinesResponse{
		RequiredVaccineDetails: vd,
	}
	return res, nil
}
