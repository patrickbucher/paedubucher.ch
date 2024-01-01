---
title: Ideen für 2024
subtitle: Womit soll ich mich beschäftigen?
author: Patrick Bucher
date: 2024-01-01T17:21:00
lang: de
---

Das neue Jahr hat angefangen, und die Gelegenheit ist gut um mir etwas Neues
vorzunehmen. Nicht einfach bloss, weil man sich Vorsätze fürs neue Jahr nimmt,
sondern weil der Jahreswechsel mir die Möglichkeit bietet, über meine
Prioritäten nachzudenken ‒ und auch gleich mit neuen Sachen anzufangen.

# Rückblick auf 2023

Auf das Jahr 2023 habe ich bereits
[andernorts](articles/2023-12-26-jahresrueckblick-2023.html) ausführlich
zurückgeblickt. Hier ist die Kurzversion davon, quasi als Andockstelle für meine
neuen Ideen. Was habe ich 2023 geleistet?

- eine Firma gegründet
- Ruby on Rails gelernt und es produktiv für ein Kundenprojekt eingesetzt
- mich mit [Nebenläufigkeit in
  Elixir](https://github.com/patrickbucher/factorizer) befasst
- mich mit [Algorithmen](https://github.com/patrickbucher/algorithms) in
  verschiedenen Sprachen befasst
- mich mit dem LAMP-Stack und Nextcloud auf Debian GNU/Linux eingearbeitet

# Ausblick auf 2024

Aus diesen Projekten sind nun folgende Ideen entstanden, mit denen ich mich 2024
gründlicher befassen möchte:

- Erlang
- Rust
- Haskell
- TypeScript
- Shell-Skripting
- Ansible
- Hugo
- Software and Mind

Diese Ideen sind nicht neu; ich habe schon einmal in alle diese Themen
hineingeschaut. Wirklich sicher fühle ich mich aber in keinem der Themen. So
wird 2024 eher ein Jahr der Vertiefung als der Neuanfänge.

## Erlang

Mit Erlang wollte ich mich schon seit dem Ende meines Informatikstudiums im
Sommer 2020 befassen. Richtig dazu gekommen bin ich nicht. Im Standardwerk von
Joe Armstrong [Programming
Erlang](https://pragprog.com/titles/jaerlang2/programming-erlang-2nd-edition/)
war ich bislang nie über Kapitel 3 hinausgekommen. Dafür habe ich mehrmals
Elixir «angelernt» ‒ und bin dabei etwas weiter gekommen als mit Erlang.

Im Sommer habe ich einige Sortieralgorithmen mit Erlang umgesetzt. Dabei ist mir
das mächtige _Pattern Matching_ sehr positiv aufgefallen. Längerfristig möchte
ich Elixir produktiv einsetzen. Um das gründlich lernen zu können, sollte ich
mich aber zuerst mit Erlang befassen. Ich möchte wissen, was Elixir wirklich
bietet ‒ und was bereits von Erlang geboten wird.

Erlang ist wohl das Gegenteil von «sexy» und «trendy»: stabil, solide ‒ und
teilweise etwas umständlich. Mittlerweile habe ich auch die Grundlagen der
funktionalen Programmierung verinnerlicht, sodass ich mich auf die Sprache und
Technologie konzentrieren kann. Der Key-Value-Store
[ETS](https://www.erlang.org/doc/man/ets.html) und die Datenbank
[Mnesia](https://www.erlang.org/doc/man/mnesia.html) sind dabei interessante
Alternativen zu bekannten Technologien wie Redis und PostgreSQL. Es muss nicht
immer vom Mainstream akzeptiert sein…

So habe ich mich in den letzten drei Tagen durch Kapitel 1-5 in _Programming
Erlang_ durchgearbeitet und möchte das Durcharbeiten dieses Buches zu meinem
neuen «Morgenthema» machen, d.h. mich täglich als erstes damit befassen, und
wenn es auch nur 15 Minuten sind. Das dürfte mich das erste Quartal 2024
beschäftigen.

Mit Erlang würde ich gerne einen kleinen verteilten Monitoring-Stack entwickeln,
womit sich virtuelle Maschinen überwachen lassen (Ressourcenauslastung,
Log-Ereignisse usw.). Erlang scheint mir die perfekte Technologie dafür zu sein.

## Rust

Auch mit Rust habe ich letzten Sommer einige wenige Algorithmen umgesetzt und
diese automatisch getestet. Durch meine Beschäftigung mit der funktionalen
Programmierung komme ich mit Rust nun einigermassen zurecht. Das nötige
Hintergrundwissen zum Memory Management fehlt mir hingegen noch. Hier habe ich
zwar einiges gelesen, aber noch lange nicht alles verstanden.

Ich sollte noch einmal das [Rust Book](https://doc.rust-lang.org/book/)
durcharbeiten ‒ und mir eigene Beispiele zu den anspruchsvollen Themen machen
(Borrowing, Smart Pointers, Concurrency). Anschliessend sollte ich es einmal in
der Praxis anwenden. Kommandozeilenwerkzeuge oder kleinere Serveranwendungen
wären passende Projekte, auch wenn mir hierzu im Moment noch die Ideen fehlen.

## Haskell

Haskell gehörte _nicht_ zu den Programmiersprachen, mit denen ich letzten Sommer
im Rahmen meines abgebrochenen Algorithmen-Projekts befasst habe. Doch
kombiniert die Sprache genau die Vorteile der dort verwendeten Sprachen: Haskell
ist ausdrucksstark, funktional und als kompilierte Sprache sehr performant.

Grundlegende Konzepte wie die IO-Monade verstehe ich mittlerweile. Trotzdem
sollte ich mich wohl parallel zur Beschäftigung mit der Sprache auch mit der
Kategorientheorie befassen.

Was Haskell-Literatur betrifft, bin ich mittlerweile sehr gut ausgerüstet:

- [Effective Haskell](https://pragprog.com/titles/rshaskell/effective-haskell/)
- [Programming in Haskell](https://www.cs.nott.ac.uk/~pszgmh/pih.html)
- [Thinking Functionally with Haskell](https://www.cambridge.org/us/universitypress/subjects/computer-science/programming-languages-and-applied-logic/thinking-functionally-haskell)
- [Algorithm Design with Haskell](https://www.cambridge.org/us/universitypress/subjects/computer-science/algorithmics-complexity-computer-algebra-and-computational-g/algorithm-design-haskell)
- [Category Theory for Programmers](https://www.blurb.com/b/9621951-category-theory-for-programmers-new-edition-hardco)

So werde ich wohl _Effective Haskell_ und _Category Theory for Programmers_
parallel durcharbeiten. _Programming in Haskell_ habe ich schon zweimal bis zur
Hälfte durchgearbeitet und könnte ich anschliessend als Repetition
durcharbeiten. _Thinking Functionally in Haskell_ und _Algorithm Design with
Haskell_ wären dann eine Vertiefung in die funktionale Programmierung.

## TypeScript

Ich habe schon mehrere halbherzige Versuche unternommen Angular zu lernen. Zwar
könnte ich dieses Web-Framework sicherlich produktiv einsetzen und damit Geld
verdienen, die Motivation fehlte mir aber bislang immer.

TypeScript gründlich zu lernen wäre einerseits schon ein grosser Schritt in
Richtung Angular. Ausserdem würde ich dabei mein bestehendes und eingerostetes
JavaScript-Wissen auffrischen ‒ und mein Wissen über Typsysteme vertiefen
könnten. Das sollte für mich genügend Motivation sein um dranzubleiben.

Mit dem ganzen TypeScript-Wissen sollte ich mir Angular dann durch das Lesen von
bestehendem Code aneignen können. (Bücher und Kurse über moderne Web-Frameworks
sind praktisch schon beim Erscheinen veraltet und darum teilweise frustrierend
beim Durcharbeiten, zumal vieles einfach nicht wie beschrieben funktioniert.)

Ob ich mich dann wirklich mit Angular befassen möchte, ist eine andere Frage.
TypeScript wird mir auf jeden Fall nicht schaden, sondern eröffnet mir auch die
Möglichkeit, mich mit anderen Web-Frameworks zu befassen. Mit [Essential
TypeScript](https://www.manning.com/books/essential-typescript-5-third-edition)
hätte ich schon das passende Buch auf Lager. (Wobei «essential» und 568
Druckseiten mir doch ein etwas mulmiges Gefühl bereiten…)

## Shell-Skripting

Zwar ist die Kommandozeile seit meinem Umstieg auf
[dwm](http://dwm.suckless.org/) im Jahr 2010 meine primäre Benutzeroberfläche
geworden. Mein Wissen beschränkt sich aber auf die alltägliche Bedienung; sobald
es um das Schreiben anspruchsvollerer Skripte geht, stehe ich recht schnell and.

So möchte ich gelegentlich mein Wissen der Unix-Shell etwas vertiefen. Da mir
Portabilität ‒ gerade im Hinblick auf OpenBSD und FreeBSD, welche _nicht_ die
Bash als Standard-Shell verwenden ‒ sehr wichtig ist, möchte ich mich aber auf
eine möglichst POSIX-kompatible Untermenge von Shell-Mechanismen beschränken.

Die Korn-Shell, die bei OpenBSD als `pdksh` zum Einsatz kommt, wäre hierzu eine
interessante Variante und würde mir auch den Wiedereinstieg in OpenBSD
ermöglichen. Mit [Learning the Korn
Shell](https://www.oreilly.com/library/view/learning-the-korn/0596001959/) steht
auch schon das passende Buch dazu schon bereit. (Es ist aus dem Jahr 2002 ‒ als
O'Reilly noch für Qualität und Nerd-Kultur stand.)

Sollte ich noch Zeit haben, könnte ich mich dann gleich noch mit der zweiten
Ausgabe von [The AWK Programming
Language](https://www.informit.com/store/awk-programming-language-9780138269722)
in AWK vertiefen. So ein gutes Nutzen-Aufwand-Verhältnis wird mir sonst keines
der 2024 angestrebten Lernprojekte bieten.

## Ansible

Seit Ende Juni habe ich nichts mehr mit Konfigurationsmanagementwerkzeugen zu
tun. Puppet vermisse ich überhaupt nicht, doch werde ich dieses Jahr wohl
endlich Zeit dafür haben um Infrastruktur für die eigene Firma aufzubauen. Und
hierzu bietet sich Ansible an.

Ansible möchte ich bereits schon im Januar im Berufsschulunterricht verwenden,
um damit die Installation von Nextcloud (zumindest teilweise) zu automatisieren.
So werde ich mir die Grundlagen davon recht pragmatisch und schnell aneignen
müssen. Eine Vertiefung erfolgt dann, wenn ich Zeit dafür und den Bedarf danach
habe.

Eine Schnellbleiche bietet mir [DevOps for the
Desperate](https://nostarch.com/devops-desperate). Vertiefen werde ich es dann
mit [Ansible for DevOps](https://leanpub.com/ansible-for-devops). («DevOps» wird
offenbar beidenorts als Marketing-Begriff missbraucht, zumal Ansible in erster
Linie ein Werkzeug für «Ops» ist.)

## Hugo

Die Umstellung von meiner eigenen Webseite auf [Hugo](https://gohugo.io/) ist
zwar vorletztes Jahr vorerst gescheitert, zumal die einfachsten Hugo-Templates
mehr bieten, als ich für meine Webseite benötige. Für eine Firmenwebseite ‒
[composed.ch](https://www.composed.ch) ist derzeit ein blosser Platzhalter ‒
könnte Hugo aber genau das Richtige sein.

So sollte ich mich endlich einmal in diese Technologie einarbeiten.
Grundsätzlich besteht bei Hugo ein «Henne-Ei-Problem»: Man benötigt zuerst ein
Template, um vernünftig mit Hugo herumspielen zu können, möchte aber
andererseits auch Templates entwickeln können, um Hugo grundlegend verstehen zu
können. So werde ich wohl versuchen, zuerst eine Seite mit einem
Standard-Template zu bauen. Anschliessend baue ich mir ein eigenes Template, bis
die damit erstellte Seite komplett funktioniert.

Da mit [Build Websites with
Hugo](https://pragprog.com/titles/bhhugo/build-websites-with-hugo/) eher
enttäuscht hat, mache ich den nächsten Versuch mit [Hugo in
Action](https://www.manning.com/books/hugo-in-action).

## Software and Mind

Auf das Mammutwerk [Software and Mind](http://softwareandmind.com/) bin ich
schon vor einigen Jahren auf der Suckless-Mailingliste aufmerksam gemacht
worden. Zwar hat mich die grundsätzliche Skepsis Mainstream-Paradigmen gegenüber
angesprochen ‒ der Gegenentwurf des Autors (ein
[COBOL-Programm](http://softwareandmind.com/#IFOP) bestehend aus einer 50'000
zeiligen Quellcodedatei) hat mich aber eher irritiert. So habe ich mich nicht
weiter mit Andrei Sorins Ideen befasst.

Aus einer Frustration mit Mainstream-Technologien habe ich das Buch zum
Jahresende aber erneut angelesen ‒ und war über die darin geäusserte (teilweise
fundamentale) Kritik regelrecht begeistert! Dass Sorin Paradigmen wie
relationale Datenbanken und strukturierte Programmierung als fehlgeleitet
kritisiert, irritiert mich zwar gewaltig, zumal ich niemals gedacht hätte,
«hinter» diese Paradigmen zurückzugehen. Andererseits wird einem gerade
hierdurch eine neue Perspektive geboten. (Die Kritik am objektorientierten
Paradigma dürfte für mich hingegen eher eine Bestätigung als eine
Infragestellung sein.)

Sorin bezeichnet die etablierten Paradigmen als unwissenschaftlich, was er
mithilfe Poppers Demarkationsbegriff ausführlich begründet. Damit ich diese
Gedanken nachvollziehen kann, muss ich mir zunächst [The Logic of Scientific
Discovery](https://www.routledge.com/The-Logic-of-Scientific-Discovery/Popper-Popper/p/book/9780415278447)
durchlesen. Erst dann kann ich mich gründlich mit _Software and Mind_ befassen.

---

Sollte ich mich 2024 tatsächlich an diesen Plan halten können, dürfte das ein
sehr lehrreiches Jahr sein!
