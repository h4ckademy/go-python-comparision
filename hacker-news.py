import urllib
from BeautifulSoup import BeautifulSoup as Soup
from beautifulsoupselect import select
from termcolor import colored


class Article():
    title = None
    link = None
    description = "No description"

    def __init__(self, title, link):
        self.title = title
        self.link = link
        self.get_meta_description()

    def get_meta_description(self):
        soup = Soup(urllib.urlopen(self.link))
        metas = select(soup, 'meta[name=description]')

        if len(metas) > 0:
            self.description = metas[0].get('content')

        print "\r fetching " + self.link


def main():
    hacker_news_url = "https://www.reddit.com/r/programming"
    soup = Soup(urllib.urlopen(hacker_news_url))
    links = select(soup, 'a.title')
    articles = []

    for link in links:
        title = link.string
        href = link.get('href')
        article = Article(title, href)
        articles.append(article)

    for article in articles:
        print " * [" + colored(article.title, 'green') + "] (" + colored(article.link, 'blue') + ")"
        print article.description


if __name__ == '__main__':
    main()
