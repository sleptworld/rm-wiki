package Model

import (
	"github.com/google/uuid"
	"net/http"
)

func Api(Status http.ConnState,apiVersion string,params map[string]string,l int,items interface{}) interface{} {

	if Status == http.StatusOK{
		res := SuccessRes{
			ApiVersion: apiVersion,
			Params:     params,
		}
		u,err := uuid.NewUUID()
		if err != nil{
			return nil
		}
		d := data{
			ID:         u.String(),
			TotalItems: l,
			Lang: params["lang"],
			Items:      items,
		}

		res.Data = d
		return res
	} else {
		res := ErrorRes{
			ApiVersion: apiVersion,
		}
		
		err := err{
			Code:    Status,
			Message: params["Message"],
			Errors:  []errs{
				{Reason: params["Reason"]},
			},
		}

		res.Error = err

		return res
	}

}
