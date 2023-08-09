#!/usr/bin/env python3

import argparse
import requests

def parse_args() -> argparse.Namespace:
    args = argparse.ArgumentParser(description="Revalidate Content")
    args.add_argument('-s', '--secret', help='Revalidation secret', required=True)
    return args.parse_args()

PATHS_TO_REVALIDATE = [
    '/courses',
    '/courses/[slug]',
    '/category',
    '/category/[slug]',
    '/blog',
    '/blog/[slug]'
]
REVALIDATE_ROUTE = 'https://hatechnolog.com/api/revalidate'

def main():
    args = parse_args()
    for to_revalidate in PATHS_TO_REVALIDATE:
        print(f'Revalidating: "{to_revalidate}"')
        response = requests.post(REVALIDATE_ROUTE, params={
            'path': to_revalidate,
            'secret': args.secret
        })
        if not response.ok:
            print(f'Got error on {to_revalidate}: {response.content}')
            break

if __name__ == '__main__':
    main()