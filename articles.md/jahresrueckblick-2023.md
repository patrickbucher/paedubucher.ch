---
title: 'Jahresrückblick 2023'
subtitle: 'Was alles passiert ist.'
author: 'Patrick Bucher'
date: 2023-12-26T12:29:00
lang: de
---

[Ende 2021](./2021-12-17-jahresrueckblick-2021.html) aus Frustration angefangen
und mit Freude abgeschlossen, habe ich auch [für
2022](./2022-12-22-jahresrueckblick-2022-und-ideen-fuer-2023.html) einen
Jahresrückblick geschrieben ‒ und diesen sogleich mit einem Ausblick auf 2023
kombiniert.

## Ausblick im Rückblick

Womit wollte ich mich ursprünglich befassen? Beruflich einerseits mit
Containern; genauer mit Container-Registries, mit Podman und mit Kubernetes.
Andererseits mit Python; genauer mit Packaging und asyncio.

Privat war ich noch täglich fleissig an SICP und wollte das auch fortführen.
Weiter nahm ich mir Elixir mit Phoenix, Rust, Perl, Clojure, Elm, Svelte, Racket
und Erlang vor. OpenBSD und FreeBSD wollte ich aussen vor lassen und mich auf
Linux konzentrieren; hier v.a. auf systemd, Netzwerke, nftables. Weiter nahm ich
mir die Beschäftigung mit «Cloud-Themen» (z.B. Datenschutz) vor.

Was ist daraus geworden? Mit Podman habe ich mich in der Firma befasst, wobei
das Projekt ins Leere gelaufen ist. Das war es dann auch mit Containern für
das Jahr 2023. Im Moment reizt mich nichts an diesem Thema. Länger als erhofft
habe ich mich mit Barman zur Sicherung von PostgreSQL-Datenbanken
auseinandergesetzt. Das war zwar lehrreich, aber auch frustrierend.

Bei Python habe ich mich kurz in asyncio eingelesen und mich dann gefragt, was
mir das überhaupt bringen soll, wenn doch andere Sprachen über viel mächtigere
Concurrency-Modelle verfügen. Mit Packaging habe ich mich kaum befasst.

Bei den Programmiersprachen habe ich mich (wie vorgenommen) mit Clojure, Elm,
Erlang, Rust und Elixir befasst ‒ hierzu gleich mehr. Phoenix, Perl, Svelte und
Racket blieben auf der Strecke. Phoenix ist aufgeschoben aber nicht aufgehoben,
Perl schaue ich gelegentlich "wegen lustig" an ‒ aber an Svelte habe ich keine
Sekunde mehr gedacht.

Mit Linux habe ich mich im Rahmen der LPIC-1-Zertifizierung befasst. Hier steht
der Austausch mit meinen Lehrerkollegen im Vordergrund und weniger der
Wissenserwerb. Was die Cloud betrifft, habe ich mich ins Thema DSGVO im Kontext
von Cloud Computing eingelesen. OpenBSD und FreeBSD sind wie geplant auf der
Strecke geblieben; den Backup-Server habe ich aber immerhin wieder in Betrieb
genommen.

Doch womit habe ich mich wirklich befasst?

## Berufliche Veränderungen

Es war im Dezember 2022 als die Idee aufkam, für einen Bekannten aus einer
früheren Anstellung Software zu entwickeln. Die Bedingung dafür war, dass wir
hierzu eine Firma gründen. Dies ging dann zu dritt bis Februar über die Bühne.
Ab April mieteten wir ein sehr preiswertes Büro an bester Lage, worauf der
offizielle Auftrag zur Entwicklug einer rollenbasierten Zugangskontrolle für die
Arbeitsplanungssoftware des Auftraggebers folgte.

Das Backend ist in Ruby on Rails geschrieben, worin ich mich möglichst effizient
einarbeiten musste. So wurde Ruby zu meiner Brotsprache, zumal ich meine andere
Stelle auf den Sommer kündigte: Ich bemerkte schnell, dass ein Tag pro Woche
nicht genügt, um das Projekt erfolgreich zu Ende zu führen. (Ab dem neuen
Schuljahr hätte dann auch dieser freie Tag noch gefehlt, und der Samstag war ja
ohnehin schon durch die Lehrerausbildung besetzt.) So habe ich meine bisher
beste Stelle nach dreieinhalb Jahren aufgegeben. Bereut habe ich es nicht, auch
wenn ich finanziell nun auf meine Reserven zurückgreifen muss.

In der Schule wurde ich im Sommer angefragt, ob ich noch für ein Modul zum Thema
Software Testing einspringen könnte. Das Thema erschien mir bei näherer
Betrachtung recht interessant, und ich steckte einige Zeit in die Vorbereitung
der Unterlagen. Die Beispiele hierzu habe ich gemäss Abmachung mit dem neu
eingestiegenen Lehrerkollegen, der die andere Hälfte des Moduls übernahm, in C#
gemacht. Dabei habe ich mich auf eine ziemlich radikale strukturierte Untermenge
von C# konzentriert und darum wenig neues über C# gelernt.

Das Cloud-Modul, das mir letztes Jahr schon eine schlechte Laune und noch mehr
graue Haare bereitete, habe ich dieses Jahr etwas angepasst, indem ich nun
verschiedene Programme anbiete: Erstens das praktisch gleiche Programm wie
letztes Jahr, zweitens die Inbetriebnahme von Nextcloud als Alternativprogramm
und drittens ein Projekt nach Wahl für diejenigen, denen die beiden geführten
Programme wenig zusagen. Der erste Teil des Moduls (Definition des
Cloud-Begriffs, Redis, S3/Minio sowie ein neuer Teil über die DSGVO im Kontext
des Cloud Computings) ist für alle Klassen der gleiche. Im zweiten Teil wird
dann entweder mit Go ein kleines Monitoring-System angepasst, Nextcloud
aufgesetzt oder eben ein Projekt nach Wahl umgesetzt.

V.a. die Beschäftigung mit dem zweiten Programm (Nextcloud in Betrieb nehmen)
hat mich recht motiviert: So habe ich mich etwas genauer mit systemd und dem
LAMP-Stack befasst. Das Programm scheint mir bei Plattformentwicklern
einigermassen anzukommen, sowie das erste Programm bei Applikationsentwicklern
erneut gut zu funktionieren scheint. Beide Programme sind je von zwei Klassen
gewählt worden. Das dritte Programm (eigenes Projekt) wurde ebenfalls von zwei
Klassen gewählt. Ob ich dieses im nächsten Jahr erneut anbieten will, werde ich
entscheiden, nachdem ich die ganzen Projekte gesehen habe. Hier macht es die
Vielfalt der Projekte schwer, inhaltlich Unterstützung bieten zu können.

## Privates Mäandrieren

Zu Beginn des Jahres habe ich mich kurz mit Elm befasst. Ich merkte jedoch
schnell, dass ich zunächst eine andere ML-artige Sprache wie OCaml oder Haskell
lernen sollte, damit ich damit ich mich bei der Beschäftigung mit Elm auf das
Web konzentrieren kann und nicht mit dem Typsystem ringen muss. (Eine solche
Sprache lernt sich einfacher auf der REPL als in einer Web-Umgebung.)

Kurz verfolgte ich die Idee von _Stock Programs_: Eine Reihe von Programmen, die
ich in jeder neuen Programmiersprache schreiben möchte, die ich lerne. Geeignete
Programme waren schnell gefunden. Systematisch damit beschäftigt habe ich mich
jedoch dann nicht mehr. Immerhin habe ich die Ideen im Repository
[stock-programs](https://github.com/patrickbucher/stock-programs) festgehalten.
Beim Erlernen der nächsten Programmiersprache könnte ich noch detailliertere
Aufgabenstellungen und erste Beispielprogramme nachliefern.

Im Frühling habe ich mich erneut an Rust herangewagt. Ich habe darin u.a. meine
[Ligatabellenberechnung](https://github.com/patrickbucher/superleague-polyglot/tree/master/superleague-rust)
nachimplementiert. Mit Rust sollte ich mich dann im Sommer noch einmal befassen,
dazu gleich mehr…

Eine lustige Bastelei ist auch das
[Soundboard](https://paedubucher.ch/soundboard/), womit ich meinem Umfeld und
auch meinen Schülern bei passender Gelegenheit auf die Nerven gehe.

[SICP](https://github.com/patrickbucher/sicp/) musste ich Mitte März gegen Ende
des dritten Kapitels aufgeben, da ich Zeit für Ruby on Rails benötigte. Nach
einem kurzen Aufflackern im Juli liess ich SICP dann liegen und habe mich erneut
mit Clojure befasst. Die Zeit mit Scheme hat sich durchaus gelohnt, zumal ich
mich in Clojure nun richtig schnell heimisch fühlte.

In den Sommerferien, wo ich im Hotelzimmer etwa eine Stunde täglich Clojure
programmierte, hatte ich auf einer Wanderung den Einfall, mich einmal mit
Algorithmen zu befassen, indem ich ein Lehrbuch zum Thema durcharbeite und die
Algorithmen in vier Sprachen umsetze: Go, Rust, Erlang und Clojure. Das
Unterfangen habe ich im Artikel [Learning
Algorithms](./2023-07-29-learning-algorithms.html) beschrieben. Ich liess das
Projekt doch recht schnell wieder fallen, was zwei Gründe hatte:

1. Das Buch _Algorithms_ war die falsche Wahl für mein Unterfangen, zumal es
   sich sehr stark auf die Methodik fokussiert und erst viel später wieder auf
   konkrete Algorithmen.
2. Ich bemerkte den ungeheuren Nutzen von Pattern Matching. Der Erlang-Code war
   extrem konzis und gut nachvollziehbar.

Die zweite Erkenntnis war dann sogleich das Ende von meiner Beschäftigung mit
Clojure. Zwar ist die Sprache meiner Meinung nach eine der schönsten überhaupt
(gerade wegen der eingebauten Datenstrukturen List, Vector, Map und Set), aber
ein mächtiges Pattern Matching hat sie leider nicht.

So beschäftigte ich mich zuerst etwas mit Haskell und dann wieder mit Elixir,
die beide über Pattern Matching verfügen. Bei Haskell bin ich im Herbst
steckengeblieben, zumal ich neben der Arbeit für die eigene Firma und dem
Unterrichten kaum Zeit dafür fand. Elixir hingegen wurde mein «Morgenthema»:
Hiermit beschäftigte ich mich ab Oktober eine halbe Stunde pro Tag, was immerhin
zu einem grösseren Concurrency-Beispiel
([factorizer](https://github.com/patrickbucher/factorizer)) und einem
ausführlichen
[Artikel](https://github.com/patrickbucher/factorizer/releases/download/v0.0.2/Concurrent-Prime-Factorization-in-Elixir.pdf)
dazu führte.

Ab Ende November war die erste halbe Stunde des Tages dann für den
[Adventskalender 2023](https://wiki.debianforum.de/Adventskalender_2023) des
Debianforums reserviert, wofür ich dieses Jahr ganze acht Türchen in Beschlag
nahm. Die Inhalte waren grösstenteils eine Jahresresteverwertung, kamen aber
sehr gut beim Publikum an. Damit ich nicht wieder manuell BBCode eintippen
musste, habe ich mich noch kurz mit Lua befasst, um einen `pandoc`-Filter zu
schreiben. Das
[Ergebnis](https://debianforum.de/forum/pastebin/?mode=view&s=42018) ist
durchaus zweckmässig und erspart mir eine Menge lästiger Handarbeit.

Für den Adventskalender beschäftigte ich mich wieder einmal mit C und Prolog.
Ersteres für ein Beispiel mit Shared Libraries, worauf mich eine Lektion der
LPIC-Schulung gebracht hat. Ausserdem implementierte ich eine
Primzahlfaktorisierung in C, die ich dann testhalber auch per CGI in Apache
eingebunden hatte.  Mit Prolog löste ich ein sogenanntes _Logical_: Ein
Logikrätsel, das ich als Anschauungsbeispiel für den Adventskalender verwendete.
(Eine Unzulänglichkeit in der Logik und beim Testen führte dann zur wohl
ausgedehntesten Diskussion im diesjährigen Adventskalender überhaupt.)

Ich schrieb sogar noch einige Zeilen PHP um den LAMP-Stack zu testen. Ich setzte
die Primzahlfaktorisierung in PHP um, womit ich für den Adventskalender
`mod_php` mit PHP-FPM verglich. Dabei kam die Frage aus, ob es performancemässig
einen signifikanten Unterschied zwischen diesen beiden Ausführungsmodi gäbe. Ich
versuchte die Frage mit einem kleinen Go-Programm namens
[request0r](https://github.com/patrickbucher/request0r) zu beantworten, womit
ich wieder einmal etwas Concurrency in Go anwenden konnte.

Was die funktionale Programmierung angeht, so fühle ich mich nun bereit für den
produktiven Einsatz. Bei Programmiersprachen, welche die Transformationen
`filter`, `map` und `reduce` in ihrer Collection-API anbieten, finde ich mich
nun auf Anhieb zurecht.

## Hobbies

Nach dem Wegzug aus dem Einfamilienhaus vor über zehn Jahren musste ich zwei
Hobbies aufgeben: Das Schlagzeugspielen und das Rudern auf dem Ergometer.
Schlagzeug und Ruderergometer musste ich aus Platzgründen verkaufen.

Im Januar habe ich mir nun wieder ein Ruderergometer gekauft ‒ und mein Training
auf den Sommer hin intensiviert: Über drei Wochen lang habe ich jeden Tag eine
Stunde auf höchstem Widerstand gerudert, ohne dass ich dabei in irgend einer
Weise Schmerzen verspürte. Mein Körpergewicht konnte ich reduzieren und mein
allgemeines Wohlergehen steigern: Ich leide wesentlich seltener unter
Spannungskopfschmerzen als früher. In den letzten Wochen bin ich dann aus
zeitlichen Gründen nur noch selten zum Rudern gekommen. Das soll sich aber
nach den Festtagen wieder ändern!

Im Spätsommer setzte ich mir das Ziel, in einer Stunde auf höchstem Widerstand
14 Kilometer zurücklegen zu können. Ich übertraf das Ziel bereits im November;
mein Rekord liegt bei etwas über 14.3 Kilometer. Für nächstes Jahr möchte ich
die 15-Kilometer-Marke anpeilen.

Im Herbst hatte ich dann die Idee mir ein E-Drum zu kaufen, zumal ich im Büro
dafür Platz habe. Nach mehrwöchigem Überlegen und Vergleichen habe ich mich dann
dazu entschieden und habe mein einst liebstes Hobby im Oktober wieder aufnehmen
können. Hierzu habe ich mir das Roland TD-27K mit Hardware von DW gekauft, u.a.
ein DW-5000-Base-Drum-Pedal, das ich früher schon hatte.

Nach fast 12 Jahren Abstinenz und nach nur wenigen Wochen unregelmässigem Üben
konnte ich praktisch wieder alles spielen, was ich früher einmal konnte. Nur an
Ausdauer, Tempo und Timing muss ich noch etwas arbeiten. Das wichtigste ist mir
aber die Freude beim Spielen als Ausgleich zur Computerarbeit.

Das Lesen und das Spazieren kamen hingegen etwas zu kurz dieses Jahr. Das sollte
ich auf nächstes Jahr etwas höher priorisieren. Immerhin bin ich von den News
abstitent geblieben, was mir viel Zeit und noch mehr Aufregung gespart hat. Der
in Leder gebundene Incerto von Nassim Taleb, den ich letztes Jahr bestellte und
nun endlich erhalten habe, dürfte mich zum Wiederlesen der vier Bände zum
Jahresanfang motivieren.
