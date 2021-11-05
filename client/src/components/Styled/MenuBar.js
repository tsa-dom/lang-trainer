import clsx from 'clsx'
import React from 'react'
import { useTranslation } from 'react-i18next'
import Button from './Button'

const MenuBar = ({ selected, setSelected, items }) => {
  const { t } = useTranslation('translation')

  const buttons = items.map(item => {
    return (
      <Button
        key={item}
        className={clsx('groups-menubar-button', selected === item && 'active')}
        text={t(item)}
        onClick={() => setSelected(item)}
      />
    )
  })

  return (
    <nav className="page-container-menubar">{buttons}</nav>
  )
}

export default MenuBar