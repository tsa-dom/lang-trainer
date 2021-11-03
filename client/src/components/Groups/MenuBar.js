import React from 'react'
import { useTranslation } from 'react-i18next'
import Button from '../Styled/Button'
import clsx from 'clsx'

const MenuBar = ({ selected, setSelected }) => {
  const { t } = useTranslation('translation')

  const handleList = () => {
    setSelected('list')
  }

  const handleAdd = () => {
    setSelected('add')
  }

  return (
    <nav className="groups-menubar">
      <Button
        className={clsx('groups-menubar-button', selected === 'list' && 'active')}
        text={t('groups-list')}
        onClick={handleList}
      />
      <Button
        className={clsx('groups-menubar-button', selected === 'add' && 'active')}
        text={t('groups-add')}
        onClick={handleAdd}
      />
    </nav>
  )
}

export default MenuBar