# Standard library
import os
import json


IMAGE_TAG = os.environ['MESSAGE_QUEUE_IMAGE_TAG']
IMAGE_NAME = f'eu.gcr.io/mimir-185212/message-queue:{IMAGE_TAG}'


ADMIN_USER = os.environ['RABBITMQ_ADMIN_USER']
ADMIN_PASSWORD = os.environ['RABBITMQ_ADMIN_PASSWORD']
ADMIN_PASSWORD_HASH = os.environ['RABBITMQ_ADMIN_PASSWORD_HASH']
NEWSRANKER_PASSWORD_HASH = os.environ['NEWSRANKER_PASSWORD_HASH']
STREAMLISTNER_PASSWORD_HASH = os.environ['STREAMLISTNER_PASSWORD_HASH']
NEWSSCRAPER_PASSWORD_HASH = os.environ['NEWSSCRAPER_PASSWORD_HASH']


def replace(original: str, key: str, value: str) -> str:
    full_key = '${' + key + '}'
    return original.replace(full_key, value)


def get_source_definitons() -> str:
    with open('source-definitions.json', 'r') as f:
        return f.read()


def get_source_conf() -> str:
    with open('source-rabbitmq.conf', 'r') as f:
        return f.read()


def create_definitions():
    defn = get_source_definitons()
    defn = replace(defn, 'admin.user', ADMIN_USER)
    defn = replace(defn, 'admin.password_hash', ADMIN_PASSWORD_HASH)
    defn = replace(defn, 'newsranker.password_hash', NEWSRANKER_PASSWORD_HASH)
    defn = replace(defn, 'streamlistener.password_hash', STREAMLISTNER_PASSWORD_HASH)
    defn = replace(defn, 'newsscraper.password_hash', NEWSSCRAPER_PASSWORD_HASH)
    print(defn)
    with open('definitions.json', 'w+') as f:
        f.write(defn)


def create_rabbitmq_conf():
    conf = get_source_conf()
    conf = replace(conf, 'admin.user', ADMIN_USER)
    conf = replace(conf, 'admin.password', ADMIN_PASSWORD)
    print(conf)
    with open('rabbitmq.conf', 'w+') as f:
        f.write(conf)


def create_rabbitmq_configuration():
    create_definitions()
    create_rabbitmq_conf()


def build_image() -> bool:
    exit_code = os.system(f'docker build -t {IMAGE_NAME} .')
    return exit_code == 0


def delete_generated_sources():
    os.remove('definitions.json')
    os.remove('rabbitmq.conf')


def main() -> None:
    create_rabbitmq_configuration()
    build_success = build_image()
    if build_success:
        print('Build succeeded')
        delete_generated_sources()
    else:
        print('Build failed')


if __name__ == '__main__':
    main()
