import i18n from 'i18next'
import { initReactI18next } from 'react-i18next'
import fi from './locales/fi/translations.json'
import en from './locales/en/translations.json'

export const resources = {
  fi: {
    translation: fi,
    language: 'fi'
  },
  en: {
    translation: en,
    language: 'en'
  }
}

i18n
  .use(initReactI18next)
  .init({
    resources,
    fallbackLng: 'fi',
    debug: false,
    interpolation: {
      escapeValue: false,
    }
  })

export default i18n