'''gather content files for oc-man-plugin'''
import http.client as http
import os
import os.path
import logging
from urllib.parse import urlparse

import yaml
from yaml import Loader


OUTPUT_DIR = 'content'

def main():
    logging.basicConfig(level=logging.INFO)
   
    if not os.path.exists(OUTPUT_DIR):
        logging.info('creating output directory')
        os.mkdir(OUTPUT_DIR)

    titles = []
    logging.info('loading topics.yaml')
    data = open('topics.yaml').read()
    topics = yaml.load(data, Loader=Loader)
    for t in topics['topics']:
        location = 'local file' if t['location'] == '' else t['location']
        logging.info(f'found topic {t["title"]}, {location}')
        if location != 'local file':
            url = urlparse(location)
            if url.scheme not in ['http', 'https']:
                logging.error(f'skipping {t["title"]}, unknown schema for {location}')
                continue
            logging.info(f'attempting to download {location}')
            connection = http.HTTPConnection if url.scheme == 'http' else http.HTTPSConnection
            connection = connection(url.netloc)
            connection.request('GET', url.path)
            response = connection.getresponse()
            if response.status != 200:
                logging.error(f'unexpected status {response.status}  downloading content for {t["title"]}.')
            body = response.read()
            outfilename = os.path.join(OUTPUT_DIR, t['title'])
            logging.info(f'writing content to {outfilename}')
            with open(outfilename, 'w') as outfile:
                outfile.write(body.decode('utf8'))
                titles.append(t['title'])
    indexfilename = os.path.join(OUTPUT_DIR, 'index.yaml')
    logging.info(f'writing index file {indexfilename}')
    with open(indexfilename, 'w') as indexfile:
        data = yaml.dump({'titles': titles})
        indexfile.write(data)
            

if __name__ == '__main__':
    main()
