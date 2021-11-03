import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import AddForm from './AddForm'
import './index.css'
import List from './List'
import MenuBar from './MenuBar'

const Groups = () => {
  const { t } = useTranslation('translation')
  const [selected, setSelected] = useState('list')

  return (
    <div className="groups-container">
      <div className="groups-header-container">
        <div className="groups-header">{t('groups')}</div>
        <hr className="groups-linebread"></hr>
        <MenuBar
          selected={selected}
          setSelected={setSelected}
        />
      </div>
      <div className="groups-body">
        {selected === 'list' &&
          <List />
        }
        {selected === 'add' &&
          <AddForm />
        }
      </div>
    </div>
  )
}

export default Groups