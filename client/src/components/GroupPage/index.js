import React, { useEffect, useState } from 'react'
import { useSelector } from 'react-redux'
import { useHistory } from 'react-router'
import MenuBar from '../Styled/MenuBar'
import Description from './Description'
import './index.css'
import Button from '../Styled/Button'
import { BiArrowBack } from 'react-icons/bi'
import List from './List'
import AddWord from './AddWord'
import { useTranslation } from 'react-i18next'

const GroupPage = () => {
  const group = useSelector(state => state.groups.selectedGroup)
  const { t } = useTranslation('translation')
  const [selected, setSelected] = useState('group-description')
  const history = useHistory()

  useEffect(() => {
    if (!group) history.push('/groups')
  }, [])

  if (!group) return <></>

  return (
    <>
      <div className="page-container-head">
        <div className="page-container-header">
          {t('group-name')} ‒ {group.name}
          <Button onClick={() => history.push('/groups')} className="page-back-button" text={<span>
            <BiArrowBack size={30} />
          </span>} />
        </div>
        <hr className="page-container-linebreak"></hr>
        <MenuBar
          selected={selected}
          setSelected={setSelected}
          items={['group-description', 'group-word-list', 'group-word-add']}
        />
      </div>
      <div className="page-container-body">
        {selected === 'group-description' && <Description description={group.description} />}
        {selected === 'group-word-list' && <List group={group} />}
        {selected === 'group-word-add' && <AddWord setSelected={setSelected} group={group} />}
      </div>
    </>
  )
}

export default GroupPage