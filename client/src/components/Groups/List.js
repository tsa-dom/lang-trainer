import React, { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
import { useDispatch, useSelector } from 'react-redux'
import {
  setSelectedGroup,
  removeGroups as remove
} from '../../features/groupSlice'
import ItemList from '../ItemList'
import { useNavigate } from 'react-router-dom'
import { removeGroups } from '../../services/groups'
import useFetch from '../../hooks/fetcher'
import { Button } from 'react-bootstrap'
import AddGroup from './AddGroup'

const List = () => {
  const groups = useSelector(state => state.groups.values)
  const { fetchGroups } = useFetch()
  const { t } = useTranslation()
  const dispatch = useDispatch()
  const navigate = useNavigate()
  const [showAddModal, setShowAddModal] = useState(false)

  useEffect(fetchGroups, [])

  const columns = [
    { field: 'name', headerName: t('groups-list-name'), width: 200 },
    { field: 'description', headerName: t('groups-list-description') },
  ]

  const handleGroupClick = (row) => {
    const words = row.words
    const group = {
      ...row,
      words: words ? words : []
    }
    dispatch(setSelectedGroup(group))
    navigate('/group')
  }

  const handleGroupRemove = async (values) => {
    if (values.length <= 0) return
    const ids = await removeGroups({ groupIds: values })
    if (ids) {
      dispatch(remove(ids))
      navigate('/groups')
    }
  }

  const handleAddGroup = () => setShowAddModal(true)

  const handleClose = () => setShowAddModal(false)

  return (
    <>
      <ItemList
        rows={groups}
        columns={columns}
        onCellClick={handleGroupClick}
        handleItemRemove={handleGroupRemove}
        title={t('groups')}
      />
      <AddGroup
        show={showAddModal}
        handleClose={handleClose}
      />
      <Button
        className='button-menu'
        onClick={handleAddGroup}
      >
        {t('add-group')}
      </Button>
    </>
  )
}

export default List