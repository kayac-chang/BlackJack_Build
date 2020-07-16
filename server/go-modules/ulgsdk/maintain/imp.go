package maintain

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/ulgsdk/ulg168utils"
)

// func GetMaintain() (*Maintain, error) {
// 	tmp := strings.Replace(ulg168utils.MaintainPath, ":game_id", ulg168utils.Conf.GameID, -1)
// 	h := &ulg168utils.Acion{
// 		Host: ulg168utils.Conf.APIHost,
// 		Path: tmp,
// 		Cond: nil,
// 	}

// 	res, err := httprequest.Get(h.RequestUrl(), ulg168utils.DefaultTimeout)
// 	if err != nil {
// 		return nil, err
// 	}

// 	obj := &Maintain{}
// 	if err := json.Unmarshal(res, obj); err != nil {
// 		return nil, err
// 	}

// 	return obj, nil
// }

// func MaintainDone() error {
// 	tmp := strings.Replace(ulg168utils.MaintainDonePath, ":game_id", ulg168utils.Conf.GameID, -1)
// 	h := &ulg168utils.Acion{
// 		Host: ulg168utils.Conf.APIHost,
// 		Path: tmp,
// 		Cond: nil,
// 	}

// 	_, err := httprequest.Put(h.RequestUrl(), ulg168utils.DefaultTimeout, nil)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func CloseGame() error {
	if ulg168utils.Conf.MaintainAPI == "" || ulg168utils.Conf.ENV == "" {
		return nil
	}

	log.Println("@@@@@@@@@@, CloseGame Start")
	req, _ := http.NewRequest(http.MethodPost, ulg168utils.Conf.MaintainAPI+ulg168utils.Conf.ENV, nil)

	req.Header.Set("PRIVATE-TOKEN", ulg168utils.Conf.MaintainToken)

	log.Println("@@@@@@@@@@, CloseGame URL: ", ulg168utils.Conf.MaintainAPI+ulg168utils.Conf.ENV)
	log.Println("@@@@@@@@@@, CloseGame PRIVATE-TOKEN: ", ulg168utils.Conf.MaintainToken)

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	r, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}

	defer r.Body.Close()
	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	log.Println("@@@@@@@@@@, CloseGame Response: ", string(response))

	if r.StatusCode != http.StatusOK && r.StatusCode != http.StatusCreated {
		return fmt.Errorf("%s", string(response))
	}
	return nil
}
