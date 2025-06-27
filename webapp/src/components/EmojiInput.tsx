import { type Dispatch, type SetStateAction, useCallback, useEffect, useRef, useState } from 'react'
import { BlockerFunction, useBlocker } from 'react-router'
import { Picker } from 'emoji-mart'
import { TelegramWebApp } from '../telegram'
import data from '@emoji-mart/data'
import { init } from 'emoji-mart'

// noinspection JSIgnoredPromiseFromCall
init(data)

type Emoji = {
  id: string
  keywords: string[]
  name: string
  native: string
  shortCodes: string
  unified: number
}

type EmojiSelector = (emoji: Emoji, event: MouseEvent) => void

export function EmojiInput(props: { value: string; setValue: Dispatch<SetStateAction<string>> }) {
  const [opened, setOpened] = useState(false)
  const close = useCallback(() => {
    setOpened(false)
  }, [])

  const blocker = useBlocker(
    useCallback<BlockerFunction>(({ historyAction }) => historyAction === 'POP' && opened, [opened]),
  )
  useEffect(() => {
    if (blocker.state === 'blocked') {
      setOpened(false)
      blocker.reset()
    }
  }, [blocker])

  return (
    <>
      <button type="button" onClick={() => setOpened((o) => !o)}>
        {props.value ?? '?'}
      </button>
      {opened && (
        <EmojiPicker
          close={close}
          onSelect={(emoji) => {
            props.setValue(emoji.native)
            setOpened(false)
          }}
        />
      )}
    </>
  )
}

function EmojiPicker(props: { close: () => void; onSelect: EmojiSelector } & Record<string, unknown>) {
  const { close, onSelect } = props
  const initialProps = useRef(props)
  const ref = useRef<HTMLDivElement>(null)
  const instance = useRef<Picker>(null)

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      const current = ref.current
      const target = event.target
      if (target instanceof Node && current && !current.contains(target)) {
        close()
      }
    }
    document.addEventListener('click', handleClickOutside, true)
    return () => {
      document.removeEventListener('click', handleClickOutside, true)
    }
  }, [close])

  useEffect(() => {
    instance.current = new Picker({
      autoFocus: true,
      emojiSize: 36,
      emojiButtonSize: 48,
      dynamicWidth: true,
      // perLine: 10,
      // i18n: {
      //   search: 'Search',
      //   notFound: 'Not found',
      // },
      onEmojiSelect: initialProps.current.onSelect,
      previewPosition: 'none',
      navPosition: 'top',
      categories: ['frequent', 'objects', 'places', 'activity', 'foods', 'symbols', 'people', 'nature', 'flags'],
      theme: TelegramWebApp.colorScheme ?? 'auto',
      data: data,
      ref,
    })
    return () => {
      instance.current = null
    }
  }, [])

  useEffect(() => {
    instance.current?.update({ onEmojiSelect: onSelect })
  }, [onSelect])

  useEffect(() => {
    const onThemeChanged = () => instance.current?.update({ theme: TelegramWebApp.colorScheme })
    TelegramWebApp.onEvent('themeChanged', onThemeChanged)
    return () => {
      TelegramWebApp.offEvent('themeChanged', onThemeChanged)
    }
  }, [])

  return (
    <>
      <div className="pt-tg-content-safe-top pb-tg-content-safe-bottom pl-tg-content-safe-left pr-tg-content-safe-right fixed top-0 right-0 bottom-0 left-0 z-20 overflow-hidden backdrop-blur-xl">
        <div className="mx-auto w-10/12 max-w-md overflow-auto overscroll-none opacity-100 *:w-full" ref={ref}></div>
      </div>
    </>
  )
}
