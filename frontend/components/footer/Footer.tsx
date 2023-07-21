import Credits from './Credits';
import Donate from './Donate';
import Social from './Social'
import { FetchLatestStats } from '@/utils/rpc/stats'

export default async function Footer() {
    const stats = await FetchLatestStats();
    return (
        <div className="w-full flex flex-col gap-4 text-center">
            <Social {...stats}/>
            <Donate/>
            <Credits {...stats}/>
        </div>
    )
}