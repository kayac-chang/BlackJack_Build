package admin

// func Validate(c *gin.Context, token string) (bool, error) {
// 	h := &ulg168utils.Acion{
// 		Host: ulg168utils.Conf.APIHost,
// 		Path: ulg168utils.AdminValidatePath,
// 		Cond: &adminValidateCond{
// 			Token: token,
// 		},
// 	}

// 	_, err := httprequest.Put(h.URL(), ulg168utils.DefaultTimeout, h.Cond.Body())
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }
