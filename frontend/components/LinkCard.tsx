import React, { useState } from 'react';
import { ILink } from '../pages';

const COPIED_TIMER = 2 * 1000;

interface IProps {
  link: ILink;
  highlightEnabled: boolean;
}

const BLUE_BUTTON_STYLES = 'text-blue-700 bg-blue-100';
const GREEN_BUTTON_STYLES = 'text-green-700 bg-green-100';

export default function LinkCard({ link, highlightEnabled }: IProps) {
  const [showCopied, setShowCopied] = useState<boolean>(false);

  const toggleShowCopied = () => {
    setShowCopied(true);
    setTimeout(() => {
      setShowCopied(false);
    }, COPIED_TIMER);
  };

  const copyLinkToClipboard = () => {
    navigator.clipboard.writeText(link.shortUrl);
    toggleShowCopied();
  };

  return (
    <div
      className={`py-4 px-8 transition-colors duration-500 lg:flex lg:items-center lg:justify-between ${
        highlightEnabled ? 'bg-blue-200' : ''
      }`}
    >
      <p className="mr-4 truncate font-light text-gray-500">
        {link.redirectUrl}
      </p>
      <div className="flex flex-col lg:flex-row lg:items-center">
        <a href={link.shortUrl} target="_blank" rel="noreferrer">
          <p className="mr-4 cursor-pointer py-4 font-light text-blue-700 transition-colors hover:text-gray-400">
            {link.shortUrl}
          </p>
        </a>
        <button
          type="button"
          className={`w-full rounded py-2 font-light transition-colors lg:w-24 ${
            showCopied ? GREEN_BUTTON_STYLES : BLUE_BUTTON_STYLES
          }`}
          onClick={copyLinkToClipboard}
        >
          {showCopied ? <span>&#10003; Copied</span> : 'Copy'}
        </button>
      </div>
    </div>
  );
}
