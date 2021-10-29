import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import './index.css'
import Menu from './Menu'
import Hamburger from 'hamburger-react'

const AppBar = ({ setPage }) => {
  const { t } = useTranslation('translation')
  const [showMenu, setShowMenu] = useState(false)

  const handleShowMenu = (value) => {
    setShowMenu(value)
  }

  return (
    <div className="appbar-container">
      <div className="appbar-logo">{t('appbar-logo')}</div>
      <Menu setPage={setPage} showMenu={showMenu} />
      <div className="appbar-burger" >
        <Hamburger toggled={showMenu} onToggle={handleShowMenu} size={80} />
      </div>
    </div>
  )
}

export default AppBar