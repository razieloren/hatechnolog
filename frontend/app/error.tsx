'use client' // Error components must be Client Components
 
import Link from 'next/link'
import { useEffect } from 'react'
 
export default function Error({
  error,
  reset,
}: {
  error: Error
  reset: () => void
}) {
  useEffect(() => {
    // Log the error to an error reporting service
    console.error(error)
  }, [error])
 
  return (
    <div className="flex flex-col gap-4 items-center justify-center text-center">
      <h2>שוד ושבר! קרתה שגיאה.</h2>
      <Link href="/">
        <button className="bg-secondary rounded-ld py-2 px-4 w-fit">
            חזרו הביתה!
        </button>
      </Link>
      <button className="bg-secondary rounded-ld py-2 px-4 w-fit"
        onClick={
          // Attempt to recover by trying to re-render the segment
          () => reset()
        }
      >
        נסו שוב
      </button>
    </div>
  )
}