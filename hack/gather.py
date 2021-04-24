'''gather content files for oc-man-plugin'''
import http.client as http
import os
import os.path
import logging
import sys
from urllib.parse import urlparse

import yaml
from yaml import Loader


OUTPUT_DIR = 'content'

def main():
    logging.basicConfig(level=logging.INFO)

    if not os.path.exists(OUTPUT_DIR):
        logging.info('creating output directory')
        os.mkdir(OUTPUT_DIR)

    logging.info('loading topics.yaml')
    titles = []
    data = open('topics.yaml').read()
    topics = yaml.load(data, Loader=Loader)

    for i, t in enumerate(topics.get('topics', [])):
        location = t.get('location')
        if location is None:
            logging.error(f'empty location field at topic {i+1}')
            sys.exit(1)
        title = t.get('title')
        if title is None:
            logging.error(f'empty title entry at topic {i+1}')
            sys.exit(1)

        logging.info(f'found topic {title}, {location}')
        url = urlparse(location)
        if url.scheme not in ['http', 'https']:
            logging.error(f'skipping {title}, unknown schema for {location} at topic {i+1}')
            continue

        logging.info(f'downloading {location}')
        connection = http.HTTPConnection if url.scheme == 'http' else http.HTTPSConnection
        connection = connection(url.netloc)
        connection.request('GET', url.path)
        response = connection.getresponse()
        if response.status != 200:
            logging.error(f'unexpected status {response.status}  downloading content for {t["title"]}.')
            sys.exit(1)
        body = response.read()

        logging.info(f'writing content for {title}')
        outfilename = os.path.join(OUTPUT_DIR, title)
        with open(outfilename, 'w') as outfile:
            outfile.write(body.decode('utf8'))
            titles.append(title)

    if len(titles) == 0:
        logging.error('zero length index file, no content')
        sys.exit(1)

    logging.info(f'writing content index file')
    indexfilename = os.path.join(OUTPUT_DIR, 'index.yaml')
    with open(indexfilename, 'w') as indexfile:
        data = yaml.dump({'titles': sorted(titles)})
        indexfile.write(data)


if __name__ == '__main__':
    main()
