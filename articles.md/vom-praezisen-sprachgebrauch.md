---
title: Vom präzisen Sprachgebrauch
subtitle: Und warum dieser in der Informatik wichtig ist
author: Patrick Bucher
date: 2021-05-09T12:30:00
lang: de
---

> Besides a mathematical inclination, an exceptionally good mastery of one's
> native tongue is the most vital asset of a competent programmer.

_Edsger W. Dijkstra_, [EWD 498](https://www.cs.utexas.edu/users/EWD/transcriptions/EWD04xx/EWD498.html)

Sprache ist für mich sehr wichtig, und zwar nicht _obwohl_, sondern _weil_ ich
Informatiker bin. Obwohl ich ohne grössere Probleme in den mathematischen
Fächern durch das Studium gekommen bin, sind meine Fähigkeiten in diesem Fach
nicht besonders stark ausgeprägt. (Meine Mathematiklehrer vom Gynmanisum können
dies sicherlich bestätigen.)

Fremdsprachen lerne ich zwar auch nicht besonders schnell, aber nach längerer
Beschäftigung mit einer Sprache kann ich in dieser doch recht präzise und
flüssig kommunizieren. Noch wichtiger finde ich, dass man sich der Unterschiede
zwischen verschiedenen Sprachen bewusst wird. Denn wortweise Übersetzungen sind
nicht nur stylistisch schlecht, sondern oft schlichtweg falsch. So lässt sich
dieser Satz aus dem Deutschen:

> Ich glaube, ich spinne.

nicht eins zu eins ins Englische übersetzen:

> I believe, I spider.

Dies gilt auch für Programmiersprachen; auch diese sollten _idiomatisch_
gebraucht werden: Was in der Programmiersprache C korrekt und idiomatisch sein
mag:

    for (int i = 0; i < n; i++) {
        sum += values[i];
    }

Sollte beispielsweise in Clojure ganz anders aussehen:

    (reduce + values)

Von jemandem, der bei Sätzen wie dem folgenden nicht leicht mit den Augen rollen
muss:

> Es macht Sinn hier ausserhalb der Box zu denken und eine Extrameile zu gehen,
> damit wir das Produkt rechtzeitig ausrollen können.

Erwarte ich auch keine stilistischen Meisterleistungen im Programmcode, gerade
wenn mehrere Programmiersprachen in einem Projekt zum Einsatz kommen.
(Mangelndes Verständnis für Sprachidiome dürfte wohl auch der Grund sein, dass
JavaScript nach vielen Jahren doch noch ein `class`-Schlüsselwort erhalten hat.)

Doch braucht man gar nicht die stilistische Ebene und die Übersetzungen zwischen
verschiedenen Idiomen heranzuziehen um sprachliche Ungereimtheiten zu finden.
Oft hapert es schon beim Gebrauch einzelner Wörter:

> ‒ Das läuft auf einem _physikalischen_ Server, nicht auf einem virtuellen.

Physik ist die Lehre der Körper, und «physikalisch» bezeichnet etwas _die Lehre
der Körper betreffendes_. Das Gegenteil eines virtuellen Servers ist ein
_physischer_ Server: ein Server, der _als Körper_ vorhanden ist. Zu einem
«physikalischen» Server könnte man sich vielleicht einen chemischen oder
biologischen Server vorstellen. (Ein _physikalischer_ Server wäre einer, auf dem
physikalische Berechnungen durchgeführt werden.) Hier dürfte es sich nicht
einmal um ein Übersetzungsproblem aus dem Englischen («physical») handeln, ist
doch der _physician_ ein Arzt (der sich mit menschlichen Körpern beschäftigt)
und der _physicist_ ein Physiker (der sich mit der Lehre von Körpern befasst) ‒
Achtung: Verwechslungsgefahr!

In englischsprachigen Vorträgen wird oft (zu oft!) das Verb «to consume»
verwendet:

> We provide services and libraries that you can _consume_.

Das Wort «konsumieren» (denn dieses Wort lässt sich auf Englisch und Deutsch
gleich verwenden) bedeutet, dass etwas _ver_braucht wird. In der Informatik ist
dies etwa im Zusammenhang mit Messaging (korrekterweise) zu lesen:

> The producer sends out messages, which are _consumed_ by multiple workers.

Die Nachrichten _müssen_ konsumiert werden, ansonsten würden sie mehrfach
verarbeitet. Einen Dienst (engl. service) oder eine Programmbibliothek (engl.
library) kann man jedoch beliebig oft verwenden, ohne dass diese dabei
verbraucht werden. (Bei einem Bezahlservice _verbraucht_ man vielleicht Tokens
oder Credits wenn man den Service _verwendet_.)

Natürlich kann man Ausdrücke wie «physikalischer Server» oder «to consume a
library» pragmatisch deuten und dabei das Gemeinte korrekt vom Gesagten
unterscheiden. Diese Toleranz führt jedoch dazu, dass die Unterschiede
zwischen verschiedenen Begriffen («physisch» und «physikalisch»; «konsumieren»
und «verwenden») verloren gehen und man dadurch weniger klar kommunizieren kann:

> ‒ The message was _consumed_ by all workers.
> - So we have a race condition in the messaging component?
> - No. The workers just use `get` instead of `pop` on the message stack. Works
>   as intended…
> - But then those messages are nod _consumed_, but only _used_.
> - So what's the difference?

Oder:

> - Wir brauchen zusätzliches IT-Budget um neue physikalische Server zu
>   beschaffen.
> - Aber wir haben doch gar kein Physik-Departement an unserer Universität!?
> - Was tut denn das zur Sache? Wir brauchen die Server um Word-Dateien zu
>   archivieren.
> - Und was hat denn das mit Physik zu tun?
> - Gar nichts, aber ich will nicht, dass die Dokumente in die Cloud gelangen.
> - Ach, sie meinen einen _physischen_ Server?
> - Natürlich, davon rede ich schon die ganze Zeit!

Diese Dialoge sind natürlich weit hergeholt und dürften ‒ hoffentlich ‒ so nie
zu vernehmen sein. Schliesslich schreibe ich diese Sprachglosse nur, um mich zu
unterhalten. (Man beachte: sprachliches Feingefühl kann eine zusätzliche Quelle
von Ärgernis _und_ Heiterkeit im Alltag sein.)
