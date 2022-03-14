import React from 'react';
import { ILink } from '../pages';
import LinkCard from './LinkCard';

interface IProps {
  links: ILink[];
  highlightEnabled: boolean;
}

export default function HistoryView({ links, highlightEnabled }: IProps) {
  return (
    <div className="w-full divide-y rounded-md bg-white shadow-md">
      {links.map((link, i) => (
        <LinkCard
          // eslint-disable-next-line react/no-array-index-key
          key={i}
          link={link}
          highlightEnabled={i === 0 && highlightEnabled}
        />
      ))}
    </div>
  );
}
