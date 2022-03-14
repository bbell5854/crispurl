import React, { SyntheticEvent, useState } from 'react';
import axios from 'axios';
import { ILink } from '../pages';
import Spinner from './Spinner';

interface IProps {
  addLink: (link: ILink) => void;
}

const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;

const VALIDATION_ERROR_MESSAGE = 'Please enter a valid url.';
const API_ERROR_MESSAGE = 'Error shortening link. Please try again later.';

export default function ShortenerForm({ addLink }: IProps) {
  const [link, setLink] = useState<string>('');
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [errorMessage, setErrorMessage] = useState<string>('');

  const handleSubmit = async (e: SyntheticEvent) => {
    e.preventDefault();

    if (link === '') {
      setErrorMessage(VALIDATION_ERROR_MESSAGE);
      return;
    }

    setErrorMessage('');
    setIsLoading(true);

    try {
      const resp = await axios.post(`${API_BASE_URL}/api/v1/redirect`, {
        redirectUrl: link,
      });

      addLink({
        shortUrl: resp.data.payload.shortUrl,
        redirectUrl: resp.data.payload.redirectUrl,
      });

      setLink('');
    } catch (error) {
      // eslint-disable-next-line no-console
      console.error(error.response.data.message);

      if (error.response.status === 422) {
        setErrorMessage(VALIDATION_ERROR_MESSAGE);
      } else {
        setErrorMessage(API_ERROR_MESSAGE);
      }
    }

    setIsLoading(false);
  };

  return (
    <div className="mb-8 w-full rounded-md bg-white p-2 shadow-md">
      <form
        className="m-6 flex flex-col md:flex-row md:items-center"
        onSubmit={handleSubmit}
      >
        <input
          className="mb-4 flex-grow border py-2 px-3 leading-tight text-gray-700 last:rounded focus:outline-none md:mb-0 md:mr-4"
          type="text"
          placeholder="Enter link url to shorten"
          value={link}
          onChange={e => setLink(e.target.value)}
        />
        <button
          className="w-full rounded bg-blue-500 py-2 px-3 text-white transition-colors hover:bg-blue-400 md:w-48"
          type="submit"
        >
          {isLoading && (
            <div className="flex items-center justify-center">
              <Spinner />
              processing...
            </div>
          )}
          {!isLoading && 'Shorten'}
        </button>
      </form>
      {errorMessage && (
        <p className="text-light m-4 w-full text-center text-red-700">
          {errorMessage}
        </p>
      )}
    </div>
  );
}
