import React, { useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import { useDispatch, useSelector } from 'react-redux'
import useGroups from '../../hooks/groups'
import { setGroups as set } from '../../features/groupSlice'

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

  return (
    <table className="groups-list">
      <thead>
        <tr>
          <th>{t('name')}</th>
          <th>{t('description')}</th>
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
  )
}

export default List