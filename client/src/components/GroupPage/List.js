import React, { useEffect } from 'react'
import useWords from '../../hooks/words'

const List = ({ group }) => {
  const { getWordsInGroup } = useWords()
  console.log(group)

  useEffect(async () => {
    const words = await getWordsInGroup({
      id: group.id
    })
    console.log(words)
  }, [])

  return (
    <></>
  )
}

export default List