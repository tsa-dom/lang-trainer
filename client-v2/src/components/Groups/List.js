import React, { useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import { useDispatch, useSelector } from 'react-redux'
import {
  setSelectedGroup,
  removeGroups as remove
} from '../../features/groupSlice'
import ItemList from '../ItemList'
import { useNavigate } from 'react-router-dom'
import { fetchGroups } from '../../utils/fetcher'
import { removeGroups } from '../../services/groups'

const List = () => {
  const groups = useSelector(state => state.groups.values)
  const { t } = useTranslation()
  const dispatch = useDispatch()
  const navigate = useNavigate()

  useEffect(fetchGroups, [])

  if (!groups) return <></>
  console.log(':><:')

  const columns = [
    { field: 'name', headerName: t('groups-list-name'), flex: 1 },
    { field: 'description', headerName: t('groups-list-description'), flex: 3 },
  ]

  const handleGroupClick = (values) => {
    const words = values.row.words
    const group = {
      ...values.row,
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

  return (
    <ItemList
      rows={groups}
      columns={columns}
      onCellClick={handleGroupClick}
      handleItemRemove={handleGroupRemove}
    />
  )
}

export default List