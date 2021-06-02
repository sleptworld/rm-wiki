package v1

//
//func checkheader(c *gin.Context) error {
//	j := Middleware.NewJWT(Config.JWTKey)
//	token := c.Request.Header.Get("token")
//
//	if token == ""{
//		return errors.New("No Token")
//	}
//	if _,err := j.ParserToken(token);err == nil{
//		return nil
//	} else {
//		return err
//	}
//}
//
//func DELETEToken(c *gin.Context) {
//
//}
//
//func GETToken(c *gin.Context) {
//
//	if err := checkheader(c);err == nil{
//		c.JSON(http.StatusOK,gin.H{
//			"msg":"OK",
//		})
//	} else {
//		c.JSON(http.StatusBadRequest,gin.H{
//			"msg":err.Error(),
//		})
//	}
//}
//
//
//func POSTToken(c *gin.Context) {
//
//	if err := checkheader(c);err == nil{
//		return
//	} else {
//		Login := Model.Login{}
//		c.BindJSON(&Login)
//		if claims,success := Model.LoginCheck(&Login,DB.Db);success{
//			j := Middleware.NewJWT(Config.JWTKey)
//			createToken, err := j.CreateToken(claims)
//
//			if err != nil {
//				return
//			} else {
//				c.Header("token",createToken)
//				c.JSON(http.StatusOK,gin.H{
//					"200":"correct",
//				})
//				return
//			}
//		} else {
//			c.JSON(http.StatusBadRequest,gin.H{
//				"400":"wrong",
//			})
//			return
//		}
//	}
//}
