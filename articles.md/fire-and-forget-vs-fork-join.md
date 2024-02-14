---
title: '«Fire and Forget» oder «Fork and Join»?'
subtitle: 'Gesellschaftsanalyse aus der Perspektive der Nebenläufigkeit '
author: 'Patrick Bucher'
date: 2022-04-30T11:35:00+0100
lang: de
---

In der nebenläufigen Programmierung ist _Fire and Forget_ ein Prinzip, bei dem
untergeordnete Aufgaben gestartet werden und anschliessend nicht mehr auf deren
Beendigung gewartet wird. Man kann davon ausgehen, dass der Task dann irgendwann
zu Ende sein wird, und das Ergebnis wird dann schon stimmen. Eine Beispiel
hierfür wäre eine Datensicherung, die von einem lang laufenden Server gestartet
wird. Sofern der Server nicht neu gestartet wird, dürfte der Vorgang dann schon
zu Ende laufen.

Beim Prinzip _Fork and Join_ wird ein Task (bzw. werden mehrere Tasks)
nebenläufig gestartet. Man wartet aber anschliessend auf ihr Ergebnis, um den
nächsten Verarbeitungsschritt erledigen zu können. Ein Beispiel hierfür wäre
etwa eine nebenläufige Sortierung eines grossen Arrays, das zunächst in
Unterarrays aufgeteilt (Fork) wird, und dessen sortierte Unterarrays dann
zusammengeführt werden (Join), wie etwa beim Merge Sort.

Beide Modelle haben ihre Berechtigung. Man sollte sich aber dessen bewusst sein,
welches Modell für eine gestellte Aufgabe das richtige ist. Und genau hier
beginnt das Problem: Nicht unbedingt ein technisches, sondern ein alltägliches:
Wie einfach ist es, mit einer neuen Sache zu beginnen, und wie schwierig ist es,
die Sache dann auch wirklich zu einem Abschluss zu bringen!

Egal ob _Fire and Forget_ oder _Fork and Join_: Es muss in beiden Fällen eine
(Unter)aufgabe gestartet werden. Beim Programmieren sind das Prozesse oder
Threads im weitesten Sinne. Im Alltag könnte das z.B. eine E-Mail sein, mit der
eine Diskussion eröffnet wird; die Zuweisung einer Aufgabe, deren Ausführung man
später kontrollieren sollte; das Anlesen eines Buches, dass dann irgendwo
herumliegt, und sich offenbar nicht von alleine fertigliest; das Erteilen einer
Hausaufgabe, für deren Kontrolle man dann doch keine Zeit hat.

Beim Programmieren ist das Prinzip _Fire and Forget_ wesentlich einfacher
umzusetzen als _Fork and Join_. Bei ersterem startet man einfach eine Aufgabe,
die dann nebenläufig oder gar parallel abläuft. Am einfachsten ist dies wohl mit
der Programmiersprache Go zu demonstrieren. Der synchrone Funktionsaufruf

    performLongRunningJob()

kann einfach durch Voranstellen des Schlüsselwort `go` in einen nebenläufigen
Funktionsaufruf verwandelt werden:

    go performLongRunningJob()

Schwieriger wird es bei Funktionen, an deren Ergebnis (bzw. Rückgabewert) man
interessiert ist. Der Funktionsaufruf

    result := performLongRunningJob()

lässt sich nicht so einfach nebenläufig ausführen; hier ist ein Channel zur
Kommunikation des Rückgabewertes nötig:

    resultChan := make(chan interface{})
    performLongRunningJob(resultChan)
    result := <-resultChan

Hier dürfte das `async`/`await`-Muster anderer Programmiersprachen syntaktisch
einfacher sein, aber eine Erkennis bleibt: _Fork and Join_ ist anspruchsvoller
als _Fire and Forget_ ‒ und für viele Anwendungsfälle der einzig zielführende
Weg.

Unsere Zivilisation erlaubt uns vielerorts das unkompliziertere _Fire and
Forget_ zu betreiben. Den Kehricht werfen wir in einen Gebührensack und diesen
dann in einen Container. Der Hauswart kümmert sich darum, dass der Container am
Sammeltag am Strassenrand steht, und die Müllabfuhr holt dann alles ab. Dahinter
verbirgt sich eine gewaltige logistische Leistung, denke man nur an den Umfang
des Entsorgungskalenders, der einem einmal jährlich per Post zugestellt wird.
(Hierfür soll es ja mittlerweile Apps geben, welche die ganze Lieferkette noch
wesentlich komplizierer machen, zumal noch IT darin involviert ist.)

Der _Forget_-Teil ist hier aber oftmals subjektiver Natur; d.h. wir denken nicht
länger an die vorabends in den Müll geworfene Chipstüte. Die Müllabfuhr und die
Kehrichtverbrennungsanlage muss sich aber noch darum kümmern. Dort setzt das
Vergessen erst ein, wenn der Abfall verbrannt ist. Die dabei ausgestossenen Gase
und der sich an der Verbrennungsanlage absetzende Russ sind dann wiederum die zu
lösenden Probleme anderer Industriezweige (etwa von Zertifikatshändlern und
Kaminfegern).  Solange aber die Müllverbrennungsanlage ordnungsgemäss gewartet 
wird, braucht man auch hier nicht mehr an den konkreten Müll zu denken: _Fire
and Forget_.

Komplizierter ist die Sache bei der Entsorgung von Elektronikartikeln oder
Fahrzeugen. Sobald diese nicht mehr leistungsfähig genug sind oder nicht mehr
den strengen Regelwerken unserer fortgeschrittenen Zivilisation
(Sicherheitsanforderungen, Strassenverkehrsordnung) entsprechen, werden sie
nicht etwa entsorgt, sondern in einen anderen Weltteil mit weniger hohen
Leistungs- und Sicherheitsanforderungen exportiert. Elektronik wird teilweise in
offenen Feuern verbrannt, um an die Kupferkabel zu gelangen.  Fahrzeuge hingegen
werden teilweise noch jahrelang betrieben, bevor sie dann zwecks Metallgewinnung
auch zerlegt werden. Zwischen _Fire_ und _Forget_ können hier also Jahrzehnte
liegen; Gesundheitsschäden durch eingeatmete Verbrennungsgase können gar ein
Leben lang nachwirken, wenn nicht sogar auf Folgegenerationen.

Wir machen es uns also oft etwas zu leicht mit dem _Fire and Forget_-Ansatz, wo
wir wieder zurück beim Programmieren wären: Sollte ein asynchron gestarteter
Backup-Vorgang nicht vielleicht doch abgewartet werden, dass man Erfolg und
Misserfolg zum Schluss zurückmelden könnte? Wird dies ausgelassen, und wird der
lang laufende, übergeordnete Serverprozess (Elternprozess) häufig (etwa aufgrund
eines Memory Leaks) neu gestartet, dürfte der lange andauernde Backup-Vorgang
(Kindprozess) nicht erfolgreich ablaufen. Bemerkt wird dies erst dann, wenn man
durch Datenverlust zu einem Restore gezwungen wird: wann es also bereits schon
zu spät ist. Erscheint uns hier _Fire and Forget_ zunächst als legitim, kommen
wir bald zur Erkenntnis, dass _Fork and Join_ wohl doch die bessere, wenn auch
schwierigere Lösung gewesen wäre.

Mein wachsender Stapel angelesener Bücher und der weniger werdende, da durch
Elektronik eingenommene Platz in meiner Wohnung weisen mich immer öfters darauf
hin, dass ich es mir mit _Fire and Forget_ wohl etwas zu einfach mache.

Bei Elektronikartikeln haben wir vor vielen Jahren die beim Kauf zu entrichtende
vorgezogene Entsorgungsgebühr eingeführt. So kann man ausgediente Geräte einfach
zur Verkaufsstelle zurückbringen, die den Artikel dann der ordnungsgemässen
Entsorgung zuführen sollte: Brennende Haufen von Elektroschrott auf
afrikansichen Müllkippen sollten dann eigentlich nicht mehr von unserer
Elektronik befeuert werden. («Sollten», denn aus Sicht des Kunden haben wir es
hier ja mit _Fire and Forget_ zu tun; ich persönlich habe die
Elektronik-Entsorgungskette noch nie versucht nachzuverfolgen.)

Bei persönlichen Interaktionen, etwa beim Vorgesetztenverhalten am Arbeitsplatz
oder beim Lehrer-Schüler-Verhältnis in Schulen, lässt sich das Problem nicht so
einfach lösen: Vorgesetzter und Lehrer können nicht beim Erteilen der Aufgabe
eine vorgezogene Kontroll- und Korrekturgebühr (gemessen in Zeiteinheiten)
entrichten, wodurch dann dieser abschliessende Vorgang garantiert durchgeführt
wird. Stattdessen muss hier auf einem virtuellen Zeitkonto eine entsprechende
Ressource in der Zukunft reserviert werden, die dann nicht von einem anderen
Vorgang eingenommen werden darf. Wir müssen also mit dem Erteilen einer Aufgabe
auch immer ein Versprechen für die Zukunft machen, um deren Erledigung
kontrollieren zu können. Hält man sich nicht an dieses Versprechen, oder gibt
man es zu Beginn gar nicht ab, verwenden wir das bequemere _Fire and Forget_ wo,
wir stattdessen _Fork and Join_ machen müssten.

So wäre ich an der These dieses Artikels angekommen: Viele Probleme der modernen
Gesellschaft ‒ Umweltverschmutzung, Zeitnot, Führungsversagen, die oft zitierte
Bildungskatastrophe ‒ sind dadurch mitverschuldet, dass man sich zu oft _Fire
and Forget_ begnügt, wo man mit _Fork and Join_ arbeiten müsste. Etwas
überspitzt könnte man den Begriff _Verantwortung_ mit dem Prinzip _Fork and
Join_ gleichsetzen: Man löst nicht nur einen Vorgang aus, sondern kümmert sich
dann auch um dessen Folgen. Oder genauer: beim Auslösen eines Vorgangs ist man
sich nicht nur dessen bewusst, dass man zu einem späteren Zeitpunkt die Folgen
desselben zu bewältigen haben wird, sondern auch, dass man für verschiedene
eingetretene Szenarien einen Plan bereithalten muss. (Hat der Schüler die
Aufgabe zufriedenstellen gelöst, wird er dafür gelobt; andernfalls fordert man
ihn zur Nacharbeit auf: ein erneuter _Fork and Join_-Vorgang wird initiiert.)

Auch hier ist ein Blick auf die Softwareentwicklung erleuchtend, zumal
gewisse Programmiersprachen bei der Auswertung von Funktionsergebnissen ein
_erschöpfendes_, d.h. alle Fälle berücksichtigendes _Pattern Matching_
erfordern; bzw. man explizit angeben muss, dass man für gewisse Ausgänge keine
Reaktion vorgesehen hat, um so von Compilerwarnungen verschont zu bleiben. Als
Programmierer muss man somit über die Probleme nachdenken, bevor sie eingetreten
sind. Zurück in den Alltag übersetzt, könnte man sich also bei einer Handlung
folgende Fragen stellen:

1. Genügt das Prinzip _Fire and Forget_, oder benötige ich einen Plan, wie ich
   mit den Konsequenzen meiner Handlung umgehen muss (_Fork and Join_)?
2. Ist mein _Pattern Matching_ erschöpfend, oder habe ich auf jetzt schon
   erwartbare Ausgänge meiner Handlung gar keinen Plan?

Die Welt ist kein Computer und die Menschen sind keine Prozesse. Dennoch bin ich
gespannt, ob diese Analogie bei der Analyse ersterer hilfreich sein wird.
