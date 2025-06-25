import {
  useDebugInformationVisibility,
  useDemoButtonsVisibility,
  useHeaderVisibility,
} from '../stores/visibilityStores'

export function ConfigForm() {
  const toggleDebug = useDebugInformationVisibility((state) => state.toggle)
  const toggleHeader = useHeaderVisibility((state) => state.toggle)
  const toggleDemoButtons = useDemoButtonsVisibility((state) => state.toggle)
  return (
    <div className="text-tg-button-text flex w-full justify-center gap-2">
      <button className="bg-tg-button rounded-l-xl p-2" onClick={toggleDemoButtons}>
        toggle demo buttons
      </button>
      <button className="bg-tg-button p-2" onClick={toggleHeader}>
        toggle header
      </button>
      <button className="bg-tg-button rounded-r-xl p-2" onClick={toggleDebug}>
        toggle debug information
      </button>
    </div>
  )
}
