import React, { useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import { useDispatch, useSelector } from 'react-redux'
import useGroups from '../../hooks/groups'
import { setGroups as set } from '../../features/groupSlice'
import { DataGrid } from '@mui/x-data-grid'

const List = () => {
  const fetched = useSelector(state => state.groups.fetched)
  const groups = useSelector(state => state.groups.values)
  const { t } = useTranslation('translation')
  const { getGroups, groups: fetchedGroups } = useGroups()
  const dispatch = useDispatch()

  useEffect(() => {
    if (!fetched) getGroups()
  }, [])

  useEffect(() => {
    if (fetchedGroups && !fetchedGroups.errors) dispatch(set(fetchedGroups))
  }, [fetchedGroups])

  const columns = [
    { field: 'name', headerName: t('groups-list-name'), width: 180 },
    { field: 'description', headerName: t('groups-list-description'), width: 150 },
  ]

  return (
    <div style={{ height: 'calc(100vh - 300px)', width: '100%' }}>
      <DataGrid
        rows={groups}
        columns={columns}
        checkboxSelection
        disableSelectionOnClick
        onCellClick={(value) => console.log(value)}
      />
    </div>
  )

  /* return (
    <table className="groups-list">
      <thead>
        <tr>
          <th>{t('groups-list-name')}</th>
          <th>{t('groups-list-description')}</th>
        </tr>
      </thead>
      <tbody>
        {groups && groups.map(group =>
          <tr key={group.id}>
            <td>{group.name}</td>
            <td>{group.description}</td>
          </tr>
        )}
      </tbody>
    </table>
  ) */
}

export default List