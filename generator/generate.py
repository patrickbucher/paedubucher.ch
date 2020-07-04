#!/usr/bin/env python3

from collections import namedtuple
import os
import re
import shutil
import sys

import markdown
from jinja2 import Environment, PackageLoader, select_autoescape
import yaml


Article = namedtuple('article', [
    'title',
    'subtitle',
    'author',
    'date',
    'lang',
    'content',
    'name',
    'href'])

env = Environment(
    loader=PackageLoader('generate', 'templates'),
    autoescape=select_autoescape(['html'])
)


def main():
    page_dir = get_page_dir()
    html_dir, article_dir = scaffold(page_dir)

    article_md_dir = os.path.join(page_dir, 'articles.md')
    articles = parse_articles(article_md_dir, article_dir)
    articles.sort(key=lambda article: (article.date, article.title))
    articles.reverse()

    index_template = env.get_template('index.jinja2')
    article_template = env.get_template('article.jinja2')

    index_content = index_template.render(articles=articles)
    index_html = os.path.join(html_dir, 'index.html')
    with open(index_html, 'w') as f:
        f.write(index_content)

    for article in articles:
        article_content = article_template.render(article=article)
        article_file = os.path.join(html_dir, article.href)
        with open(article_file, 'w') as f:
            f.write(article_content)


def scaffold(page_dir):
    html_dir = os.path.join(page_dir, 'html')
    if os.path.exists(html_dir) and os.path.isdir(html_dir):
        shutil.rmtree(html_dir)
    os.mkdir(html_dir)

    article_dir = os.path.join(html_dir, 'articles')
    os.mkdir(article_dir)
    style_from = os.path.join(page_dir, 'style', 'style.css')
    style_to = os.path.join(html_dir, 'style.css')
    shutil.copyfile(style_from, style_to)

    return html_dir, article_dir


def parse_articles(article_md_dir, article_dir):
    articles = []
    for article_file in os.listdir(article_md_dir):
        article_path = os.path.join(article_md_dir, article_file)
        meta_data, html_content = parse_article(article_path)
        date_str = meta_data['date'].strftime('%Y-%m-%d')
        title = meta_data['title'].strip()
        subtitle = meta_data['subtitle'] if 'subtitle' in meta_data else ''
        name = normalize(f'{date_str}-{title}')
        article = Article(
            title=meta_data['title'],
            subtitle=subtitle.strip(),
            author=meta_data['author'],
            date=date_str,
            lang=meta_data['lang'],
            content=html_content,
            name=name,
            href=os.path.join('articles', f'{name}.html')
        )
        articles.append(article)
    return articles


def normalize(s):
    s = re.sub(r'\s', '-', s)
    s = re.sub('[^-_0-9a-zA-Z]+', '', s)
    s = s.lower()
    return s


def parse_article(article_path):
    with open(article_path) as f:
        content = f.read()

    sep = '---'
    meta_start = content.index(sep)
    if meta_start == -1:
        raise ValueError(f'{article_path}: missing metadata start')
    meta_end = content.index(sep, meta_start+len(sep))
    if meta_end == -1:
        raise ValueError(f'{article_path}: missing metadata end')
    meta_data = content[meta_start+len(sep)+1:meta_end]

    yaml_dict = yaml.load(meta_data, Loader=yaml.SafeLoader)
    content_data = content[meta_end+len(sep)+1:]
    html = markdown.markdown(content_data, output_format='html5')

    return yaml_dict, html


def get_page_dir():
    if len(sys.argv) < 2:
        raise ValueError('missing directory argument')

    page_dir = sys.argv[1]
    if not os.path.exists(page_dir):
        raise ValueError(f'the directory {page_dir} does not exist')

    return page_dir


if __name__ == '__main__':
    main()
