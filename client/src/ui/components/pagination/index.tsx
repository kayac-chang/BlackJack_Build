import { useState } from 'react';
import { clamp } from 'ramda';

export default function usePagination<T>(data: T[], itemsPerPage: number) {
  const min = 1;
  const max = Math.ceil(data.length / itemsPerPage);

  const [current, setCurrent] = useState(min);
  const range = clamp(min, max);

  function next() {
    setCurrent((current) => range(current + 1));
  }

  function prev() {
    setCurrent((current) => range(current - 1));
  }

  function jump(page: number) {
    setCurrent(range(page));
  }

  return {
    get data() {
      const begin = (current - 1) * itemsPerPage;
      const end = begin + itemsPerPage;
      return data.slice(begin, end);
    },

    get page() {
      return current;
    },

    get range() {
      return { max, min };
    },

    next,
    prev,
    jump,
  };
}
