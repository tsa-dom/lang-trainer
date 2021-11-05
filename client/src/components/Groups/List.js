import React, { useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import { useDispatch, useSelector } from 'react-redux'
import useGroups from '../../hooks/groups'
import { setGroups as set, setSelectedGroup } from '../../features/groupSlice'
import { DataGrid } from '@mui/x-data-grid'
import { useHistory } from 'react-router'

const List = () => {
  const fetched = useSelector(state => state.groups.fetched)
  const groups = useSelector(state => state.groups.values)
  const { t } = useTranslation('translation')
  const { getGroups } = useGroups()
  const dispatch = useDispatch()
  const history = useHistory()

  useEffect(async () => {
    if (!fetched) {
      const groups = await getGroups()
      dispatch(set(groups))
    }
  }, [])

  const columns = [
    { field: 'name', headerName: t('groups-list-name'), width: 180 },
    { field: 'description', headerName: t('groups-list-description'), width: 150 },
  ]

  const handleGroupClick = (values) => {
    dispatch(setSelectedGroup(values.row))
    history.push('/group')
  }

  return (
    <div style={{ height: 'calc(100vh - 300px)', width: '100%' }}>
      <DataGrid
        rows={groups}
        columns={columns}
        checkboxSelection
        disableSelectionOnClick
        onCellClick={handleGroupClick}
      />
    </div>
  )
}

export default List