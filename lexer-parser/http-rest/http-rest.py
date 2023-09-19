#!/usr/bin/python3

import requests


class http_rest:
    def __init__(self, url):
        self.url = url
    

    def query(self,input):
        print('query server')

        payload = self.query_to_json(input)

        try:

            response = requests.get(self.url, payload)

            if response.status_code == 200:
                print('it worked')
            else:
                print('it didnt work')

        except Exception as e:
            print("There was an error attepting to query the server!")



    def query_to_json(self, input):
        print('converting to json')
        payload = {
            "statement" : str(input)
        }

        return payload

if __name__ == "__main__":
    print("http-rest !!")

    hr = http_rest("http://localhost:8080/datatwo")

    hr.query("MATCH (a:bike) RETURN a")
