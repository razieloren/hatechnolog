// POST: https://<HOST>/api/revalidate?path=<PATH>&secret=<TOKEN>

import { NextRequest, NextResponse } from 'next/server'
import { revalidatePath } from 'next/cache'
 
export async function POST(request: NextRequest) {
  const secret = request.nextUrl.searchParams.get('secret')
  const path = request.nextUrl.searchParams.get('path')
 
  if (secret !== process.env.REVALIDATE_TOKEN) {
    return NextResponse.json({ message: 'Unauthorized' }, { status: 401 })
  }
  if (!path) {
    return NextResponse.json({ message: 'Bad request' }, { status: 400 })
  }
 
  revalidatePath(path)
 
  return NextResponse.json({ revalidated: true, now: Date.now() })
}