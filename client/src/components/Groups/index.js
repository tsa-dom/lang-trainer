import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import AddForm from './AddForm'
import './index.css'
import List from './List'
import MenuBar from '../Styled/MenuBar'

const Groups = () => {
  const { t } = useTranslation('translation')
  const [selected, setSelected] = useState('groups-list')

  return (
    <>
      <div className="page-container-head">
        <div className="page-container-header">{t('groups')}</div>
        <hr className="page-container-linebreak"></hr>
        <MenuBar
          selected={selected}
          setSelected={setSelected}
          items={['groups-list', 'groups-add']}
        />
      </div>
      <div className="page-container-body">
        {selected === 'groups-list' &&
          <List />
        }
        {selected === 'groups-add' &&
          <AddForm setSelected={setSelected} />
        }
      </div>
    </>
  )
}

export default Groups