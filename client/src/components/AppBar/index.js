import React from 'react'
import { useTranslation } from 'react-i18next'
import { useHistory } from 'react-router'
import './index.css'
import Menu from './Menu'

const AppBar = () => {
  const { t } = useTranslation('translation')
  const history = useHistory()

  const toMainPage = () => {
    history.push('/')
  }

  return (
    <div className="appbar-container">
      <div className="appbar-logo" onClick={toMainPage}>{t('appbar-logo')}</div>
      <Menu
      />
    </div>
  )
}

export default AppBar