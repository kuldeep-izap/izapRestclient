package izapRestclient

import (
		"net/http"
		"io/ioutil"
		
)

type InputStruct struct{
  Url string
  Headers map[string]string
}

func GetResponse(input InputStruct) ([]byte,  error){
	
	req,err := http.NewRequest("GET", input.Url, nil)
	for name,val := range input.Headers{
		req.Header.Set(name,val)
	}
	
	if err != nil {
        return nil,err
	}
	
	// Send the request via a client
	client := &http.Client{}
	resp,err := client.Do(req)
	if err != nil {
    	return nil, err
	}
	
	defer resp.Body.Close()
    // Read the content into a byte array
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
       return nil, err
    }
	
	return body, nil
}

func GetHeaders(input InputStruct)(http.Header,error){
	req,err := http.NewRequest("GET", input.Url, nil)
	for name,val := range input.Headers{
		req.Header.Set(name,val)
	}
	
	if err != nil {
        return nil,err
	}
	
	// Send the request via a client
	client := &http.Client{}
	resp,err := client.Do(req)
	if err != nil {
    	return nil, err
	}
	
	return resp.Header, nil
}



		
