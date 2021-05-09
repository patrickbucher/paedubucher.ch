#!/usr/bin/env python3

from collections import namedtuple
from datetime import datetime
import os
import re
import shutil
import sys
import xml.etree.ElementTree as ET
from xml.dom import minidom

import markdown
from jinja2 import Environment, PackageLoader, select_autoescape
import yaml


Article = namedtuple('article', [
    'title',
    'subtitle',
    'author',
    'datetime_raw',
    'date_raw',
    'date_fmt',
    'lang',
    'content',
    'markdown',
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
    articles.sort(key=lambda article: (article.datetime_raw, article.title))
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

    atom_feed_content = create_atom_feed(articles)
    atom_feed_xml = os.path.join(html_dir, 'atom.xml')
    with open(atom_feed_xml, 'wb') as f:
        f.write(atom_feed_content)


def scaffold(page_dir):
    html_dir = os.path.join(page_dir, 'html')
    if os.path.exists(html_dir) and os.path.isdir(html_dir):
        shutil.rmtree(html_dir)
    os.mkdir(html_dir)

    article_dir = os.path.join(html_dir, 'articles')
    os.mkdir(article_dir)

    static_dir = os.path.join(page_dir, 'static')

    for entry in ['style.css', 'favicon.ico', 'robots.txt']:
        source = os.path.join(static_dir, entry)
        target = os.path.join(html_dir, entry)
        shutil.copyfile(source, target)

    return html_dir, article_dir


def parse_articles(article_md_dir, article_dir):
    articles = []
    for article_file in os.listdir(article_md_dir):
        article_path = os.path.join(article_md_dir, article_file)
        meta_data, html_content, markdown_content = parse_article(article_path)
        lang = meta_data['lang']
        datetime_raw = meta_data['date']
        date_raw = datetime_raw.date()
        date_fmt = format_date(date_raw, lang)
        title = meta_data['title'].strip()
        subtitle = meta_data['subtitle'] if 'subtitle' in meta_data else ''
        name = normalize(f'{date_raw}-{title}')
        article = Article(
            title=meta_data['title'],
            subtitle=subtitle.strip(),
            author=meta_data['author'],
            datetime_raw=datetime_raw,
            date_raw=date_raw,
            date_fmt=date_fmt,
            lang=meta_data['lang'],
            content=html_content,
            markdown=markdown_content,
            name=name,
            href=os.path.join('articles', f'{name}.html')
        )
        articles.append(article)
    return articles


def format_date(date, lang):
    lang_formats = {
        'en': '%Y-%m-%d',
        'de': '%d.%m.%Y',
    }
    if lang not in lang_formats:
        raise ValueError(f'unknown language "{lang}"')
    return date.strftime(lang_formats[lang])


def normalize(s):
    s = re.sub(r'\s', '-', s)
    s = re.sub(u'ä', 'ae', s)
    s = re.sub(u'ö', 'oe', s)
    s = re.sub(u'ü', 'ue', s)
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

    return yaml_dict, html, content_data


def get_page_dir():
    if len(sys.argv) < 2:
        raise ValueError('missing directory argument')

    page_dir = sys.argv[1]
    if not os.path.exists(page_dir):
        raise ValueError(f'the directory {page_dir} does not exist')

    return page_dir


def create_atom_feed(articles):
    feed = ET.Element('feed')
    feed.attrib['xmlns'] = 'http://www.w3.org/2005/Atom'

    title = ET.SubElement(feed, 'title')
    title.text = 'paedubucher.ch'

    subtitle = ET.SubElement(feed, 'subtitle')
    subtitle.text = 'paedubucher.ch Article Feed'

    selfLink = ET.SubElement(feed, 'link')
    selfLink.attrib['href'] = 'https://paedubucher.ch/atom.xml'
    selfLink.attrib['rel'] = 'self'

    link = ET.SubElement(feed, 'link')
    link.attrib['href'] = 'https://paedubucher.ch/'

    identifier = ET.SubElement(feed, 'id')
    identifier.text = 'https://paedubucher.ch/'

    updated = ET.SubElement(feed, 'updated')
    updated.text = datetime.utcnow().isoformat('T') + 'Z'

    for article in articles:
        add_atom_feed_entry(article, feed)

    raw = ET.tostring(feed, 'utf-8')
    return minidom.parseString(raw).toprettyxml(indent=4*' ').encode('utf-8')


def add_atom_feed_entry(article, parent):
    entry = ET.SubElement(parent, 'entry')

    title = ET.SubElement(entry, 'title')
    title.text = article.title

    link = ET.SubElement(entry, 'link')
    href = f'https://paedubucher.ch/{article.href}'
    link.attrib['href'] = href

    author = ET.SubElement(entry, 'author')
    name = ET.SubElement(author, 'name')
    name.text = article.author

    identifier = ET.SubElement(entry, 'id')
    identifier.text = href

    updated = ET.SubElement(entry, 'updated')
    updated.text = article.datetime_raw.isoformat('T') + 'Z'

    content = ET.SubElement(entry, 'content')
    content.attrib['type'] = 'text/markdown; charset=UTF-8'
    content.text = article.markdown


if __name__ == '__main__':
    main()
