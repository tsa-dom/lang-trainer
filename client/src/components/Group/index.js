import React, { useEffect, useState } from 'react'
import { Container } from 'react-bootstrap'
import { useSelector } from 'react-redux'
import { useNavigate } from 'react-router-dom'
import WordList from './WordList'
import ModifyWord from './WordModal/ModifyWord'

const GroupPage = () => {
  const group = useSelector(state => state.groups.selectedGroup)
  const [word, setWord] = useState(undefined)
  const navigate = useNavigate()

  useEffect(() => {
    if (!group) navigate('/groups')
  }, [group])

  if (!group) return <></>

  return (
    <>
      <Container>
        <WordList
          group={group}
          word={word}
          setWord={setWord}
        />
        {word &&
          <ModifyWord word={word} setWord={setWord} />
        }
      </Container>
    </>
  )
}

export default GroupPage