import i18next from 'i18next';
import Backend from 'i18next-http-backend';
import { initReactI18next } from 'react-i18next';

type Props = {
  loadPath?: string;
};

export async function init(props?: Props) {
  //
  return await i18next
    //
    .use(initReactI18next)
    .use(Backend)
    .init({
      lng: 'en',
      fallbackLng: 'en',
      backend: {
        loadPath: props?.loadPath || `${process.env.PUBLIC_URL}/locales/{{lng}}/{{ns}}.json`,
      },
    });
}
