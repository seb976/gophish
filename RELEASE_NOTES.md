# Notes de version

## Version 0.12.1-custom.2 — Pages de sensibilisation

### Nouveauté : les pages de sensibilisation (« Training Pages »)

Vous pouvez désormais montrer un contenu pédagogique à une personne **juste après qu'elle s'est fait piéger** par une campagne de test, plutôt que de la renvoyer simplement vers un autre site.

**Ce qui change pour vous :**

- **Un nouveau menu « Training Pages »** dans la barre de navigation. Il vous permet de créer, modifier et supprimer vos pages de sensibilisation.
- **Création simple d'une page** : donnez-lui un nom, puis rédigez son contenu dans l'éditeur, ou importez une page existante depuis un site web avec le bouton « Import Site ».
- **Personnalisation** : comme pour les e-mails et les pages d'atterrissage, vous pouvez insérer le prénom, le nom ou l'adresse de la personne pour rendre le message plus parlant.
- **Liaison avec les pages d'atterrissage** : lorsque vous modifiez une page d'atterrissage avec la capture des données activée, une nouvelle liste **« Training page »** apparaît. Choisissez la page de sensibilisation à afficher après l'envoi du formulaire.
- **Priorité sur la redirection** : si une page de sensibilisation est choisie, c'est elle qui s'affiche, et non l'adresse de redirection. Choisissez « None » pour garder la redirection habituelle.

**Exemple d'usage :** une personne saisit ses identifiants sur une fausse page de connexion. Au lieu d'être redirigée sans comprendre, elle voit tout de suite une page qui lui explique que c'était un exercice, quels signes auraient dû l'alerter et comment réagir la prochaine fois.

---

## Version 0.12.1-custom.1 — Signalement des e-mails avec Microsoft 365

### Nouveauté : connexion à la boîte de signalement Microsoft 365 sans mot de passe

Gophish peut lire la boîte mail où vos collaborateurs signalent les e-mails suspects, et comptabiliser ces signalements dans les résultats de campagne. Cette fonction marche désormais avec **Microsoft 365 / Outlook**, qui n'accepte plus la connexion par simple mot de passe.

**Ce qui change pour vous :**

- Dans **Settings → Reporting Settings**, une nouvelle option **« Use OAuth2 (Microsoft 365) »** est disponible.
- Une fois cochée, trois champs apparaissent : **Tenant ID**, **Client ID** et **Client Secret**. Votre administrateur Microsoft 365 vous les fournit.
- Le champ mot de passe n'est plus nécessaire dans ce mode.
- Le bouton de test de connexion fonctionne aussi avec ce nouveau mode, pour vérifier les réglages avant de les enregistrer.

**Bénéfice :** les signalements faits par vos collaborateurs depuis une boîte Microsoft 365 remontent de nouveau automatiquement dans Gophish. Vous mesurez donc non seulement qui s'est fait piéger, mais aussi qui a eu le bon réflexe de signaler.
