import React from 'react'
import { useTranslation } from 'react-i18next'

const Templates = () => {
  const { t } = useTranslation()

  return (
    <div className="words-body">
      <div style={{ fontWeight: 'bold' }} className="words-header">{t('template-info')}</div>
    </div>
  )
}

export default Templates