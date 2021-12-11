import React, { useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { useHistory } from 'react-router'
import MenuBar from '../Styled/MenuBar'
import Description from './Description'
import Button from '../Styled/Button'
import { BiArrowBack } from 'react-icons/bi'
import List from './List'
import AddWord from './AddWord'
import { useTranslation } from 'react-i18next'
import { modifyGroup as modify } from '../../features/groupSlice'
import ModifyWord from './ModifyWord'
import Temlates from './Templates'
import { modifyGroup } from '../../services/groups'

const GroupName = ({ group }) => {
  const [editMode, setEditMode] = useState(false)
  const [value, setValue] = useState(group.name)
  const dispatch = useDispatch()

  const handleDoubleClick = () => {
    setEditMode(true)
  }

  const handleEnter = async () => {
    const res = await modifyGroup({
      ...group,
      name: value
    })
    dispatch(modify(res))
    setEditMode(false)
  }

  return (
    <>
      {!editMode &&
        <span style={{ cursor: 'pointer' }} onDoubleClick={handleDoubleClick}>
          { group.name }
        </span>
      }
      {editMode &&
        <span>
          <input style={{
            backgroundColor: 'rgb(202, 203, 209)',
            fontSize: 32
          }}
          onKeyPress={(e) => {
            if (e.code === 'Enter') handleEnter()
          }}
          value={value}
          onChange={e => setValue(e.target.value)}
          />
        </span>
      }
    </>
  )
}

const GroupPage = () => {
  const group = useSelector(state => state.groups.selectedGroup)
  const { t } = useTranslation()
  const [selectedPage, setSelectedPage] = useState('group-description')
  const [selectedWord, setSelectedWord] = useState(null)
  const history = useHistory()

  const items = ['group-description', 'group-word-list', 'group-word-add', 'word-templates']

  useEffect(() => {
    if (!group) history.push('/groups')
  }, [])

  if (!group) return <></>

  return (
    <>
      <div className="page-container-head">
        <div className="page-container-header">
          {t('group-name')} ‒ <GroupName group={group} />
          <Button onClick={() => history.push('/groups')} className="page-back-button" text={<span>
            <BiArrowBack style={{ position: 'absolute', marginTop: -5, marginLeft: -20 }} size={30} />
          </span>} />
        </div>
        <hr className="page-container-linebreak"></hr>
        <MenuBar
          selected={selectedPage}
          setSelected={setSelectedPage}
          items={items.includes(selectedPage) ? items : items.concat(selectedPage)}
        />
      </div>
      <div className="page-container-body">
        {selectedPage === 'group-description' && <Description group={group} />}
        {selectedPage === 'group-word-list' &&
          <List
            group={group}
            setSelectedWord={setSelectedWord}
            setSelectedPage={setSelectedPage}
          />
        }
        {selectedPage === 'group-word-add' && <AddWord setSelected={setSelectedPage} group={group} />}
        {!items.includes(selectedPage) &&
          <ModifyWord word={selectedWord} setSelected={setSelectedPage} />
        }
        {selectedPage === 'word-templates' && <Temlates /> }
      </div>
    </>
  )
}

export default GroupPage