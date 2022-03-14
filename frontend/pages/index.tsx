import { useEffect, useState } from 'react';
import type { NextPage } from 'next';
import Head from 'next/head';
import ShortenerForm from '../components/ShortenerForm';
import HistoryView from '../components/HistoryView';

const LOCAL_STORAGE_KEY = 'linkCache';
const HIGHLIGHT_TIMER = 2 * 1000;

export interface ILink {
  shortUrl: string;
  redirectUrl: string;
}

const Home: NextPage = () => {
  const [links, setLinks] = useState<ILink[]>([]);
  const [highlightEnabled, setHighlightEnabled] = useState<boolean>(false);

  useEffect(() => {
    const cachedLinks = localStorage.getItem(LOCAL_STORAGE_KEY);

    if (cachedLinks) {
      setLinks(JSON.parse(cachedLinks));
    }
  }, []);

  const toggleHighlight = () => {
    setHighlightEnabled(true);

    setTimeout(() => {
      setHighlightEnabled(false);
    }, HIGHLIGHT_TIMER);
  };

  const addLink = (link: ILink) => {
    const linksCopy = JSON.parse(JSON.stringify(links));

    linksCopy.unshift(link);

    setLinks(linksCopy);
    toggleHighlight();
    localStorage.setItem(LOCAL_STORAGE_KEY, JSON.stringify(linksCopy));
  };

  return (
    <div className="">
      <Head>
        <title>Crisp Url</title>
        <meta name="description" content="Url Shortner" />
      </Head>
      <div className="flex h-screen w-screen flex-col items-center bg-gray-100">
        <div className="mt-16 w-3/4">
          <div className="mb-12 flex justify-between">
            <div>
              <h1 className="font-lilita-one text-4xl text-blue-500">
                CRISP URL
              </h1>
              <h2 className="text-xl text-blue-500">
                Free online URL shortner.
              </h2>
              <h3 className="font-light text-blue-500">
                Unlimited use. No Fees. No BS.
              </h3>
            </div>
            <div className="flex flex-col-reverse">
              <a
                className="rounded shadow-md"
                style={{
                  // This is a bugfix for character decenders expanding the element height
                  lineHeight: 0,
                }}
                href="https://www.buymeacoffee.com/bbell5854"
                target="_blank"
                rel="noreferrer"
              >
                <img
                  src="https://cdn.buymeacoffee.com/buttons/default-white.png"
                  className="w-40 rounded align-top"
                  alt="Buy Me A Coffee"
                />
              </a>
            </div>
          </div>
          <div className="w-full">
            <ShortenerForm addLink={addLink} />
            {!!links.length && (
              <HistoryView links={links} highlightEnabled={highlightEnabled} />
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Home;
